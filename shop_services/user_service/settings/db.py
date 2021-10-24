import json

import nacos
from loguru import logger
from playhouse.pool import PooledMySQLDatabase
from playhouse.shortcuts import ReconnectMixin


class ReconnectMysqlDatabase(ReconnectMixin, PooledMySQLDatabase):
    pass


def update_cfg(args):
    print("配置产生变化")
    print(args)


NACOS = {
    "Host": "localhost",
    "Port": 8848,
    "NameSpace": "77a22bd7-4ce5-4af5-89f0-b427b6d6c55f",
    "User": "nacos",
    "Password": "nacos",
    "DataId": "user_svs",
    "Group": "dev"
}

MYSQL_DB = "shop_user"
MYSQL_HOST = "localhost"
MYSQL_PORT = 3306
MYSQL_USER = "root"
MYSQL_PASSWD = "root"

client = nacos.NacosClient(f'{NACOS["Host"]}:{NACOS["Port"]}',
                           namespace=NACOS["NameSpace"],
                           username=NACOS["User"],
                           password=NACOS["Password"])

# get config
data = client.get_config(NACOS["DataId"], NACOS["Group"])
data = json.loads(data)
logger.info(data)

# MySQL
DB = ReconnectMysqlDatabase(data["mysql"]["db"], host=data["mysql"]["host"], port=data["mysql"]["port"],
                            user=data["mysql"]["user"], password=data["mysql"]["password"])

# Consul
CONSUL_HOST = data["consul"]["host"]
CONSUL_PORT = data["consul"]["port"]

SERVICE_NAME = data["name"]
SERVICE_TAGS = data["tags"]
