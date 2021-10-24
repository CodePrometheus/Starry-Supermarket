import argparse
import logging
import os
import signal
import socket
import sys
import uuid
from concurrent import futures
from functools import partial

import grpc
from loguru import logger

from shop_services.common_service.grpc_health.v1 import health
from shop_services.common_service.grpc_health.v1 import health_pb2_grpc
from shop_services.common_service.register import register
from shop_services.user_service.handler.user import UserServicer
from shop_services.user_service.proto import user_pb2_grpc
from shop_services.user_service.settings import db

BASE_DIR = os.path.dirname(os.path.abspath(os.path.dirname(__file__)))
sys.path.insert(0, BASE_DIR)


def start():
    # 解析入参
    parser = argparse.ArgumentParser()
    parser.add_argument('--ip',
                        nargs="?",
                        type=str,
                        default="localhost",
                        help="binding ip"
                        )
    parser.add_argument('--port',
                        nargs="?",
                        type=int,
                        default=9000,
                        help="the listening port"
                        )
    args = parser.parse_args()

    if args.port == 0:
        port = get_free_tcp_port()
    else:
        port = args.port

    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))

    # 注册用户服务
    user_pb2_grpc.add_UserServicer_to_server(UserServicer(), server)

    # 注册健康检查
    health_pb2_grpc.add_HealthServicer_to_server(health.HealthServicer(), server)

    server.add_insecure_port(f"{args.ip}:{args.port}")

    logger.add("logs/user_service_{time}.log", rotation="2 MB", retention="1 hours")

    """
    主进程退出信号监听
        windows下支持的信号是有限的：
        SIGINT ctrl+C终端
        SIGTERM kill发出的软件终止
    """
    service_id = str(uuid.uuid1())
    signal.signal(signal.SIGINT, partial(handler_exit, service_id=service_id))
    signal.signal(signal.SIGTERM, partial(handler_exit, service_id=service_id))

    # start server
    logger.info(f"启动服务: {args.ip} : {port}")
    server.start()

    # consul register
    logger.info(f"服务注册开始")
    consul = register.ConsulRegister(db.CONSUL_HOST, db.CONSUL_PORT)
    if not consul.register(name=db.SERVICE_NAME,
                           id=service_id,
                           addr=args.ip,
                           port=args.port,
                           tags=db.SERVICE_TAGS,
                           check=None):
        logger.error(f"服务 {db.SERVICE_NAME} 注册失败")
        sys.exit(0)

    logger.info(f"服务 {db.SERVICE_NAME} 注册成功")
    server.wait_for_termination()


def handler_exit(signo, frame, service_id):
    logging.info("signo : ", signo, " frame : ", frame)
    consul = register.ConsulRegister(db.CONSUL_HOST, db.CONSUL_PORT)
    logger.info(f"注销 {service_id} 服务")
    consul.deregister(service_id=service_id)
    logger.info(f"注销 {service_id} 服务成功")
    sys.exit(0)


def get_free_tcp_port():
    tcp = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    tcp.bind(("", 0))
    _, port = tcp.getsockname()
    tcp.close()
    return port


if __name__ == '__main__':
    logging.basicConfig()
    logging.warning("==注意，我启动了 ;)")
    db.client.add_config_watcher(db.NACOS["DataId"], db.NACOS["Group"], db.update_cfg)
    start()
