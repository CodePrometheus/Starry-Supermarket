import json
from abc import ABC

import nacos
from playhouse.pool import PooledMySQLDatabase
from playhouse.shortcuts import ReconnectMixin


class ReconnectMysqlDatabase(ReconnectMixin, PooledMySQLDatabase, ABC):
    pass


def update_cfg(args):
    print("配置产生变化")
    print(args)


NACOS = {
    "Host": "localhost",
    "Port": 8848,
    "NameSpace": "52b1f242-d587-4544-8f51-2f9a4e5f6668",
    "User": "nacos",
    "Password": "nacos",
    "DataId": "goods_svs",
    "Group": "dev"
}

client = nacos.NacosClient(f'{NACOS["Host"]}:{NACOS["Port"]}',
                           namespace=NACOS["NameSpace"],
                           username=NACOS["User"],
                           password=NACOS["Password"])

# get config
data = client.get_config(NACOS["DataId"], NACOS["Group"])
data = json.loads(data)

# MySQL
DB = ReconnectMysqlDatabase(data["mysql"]["db"], host=data["mysql"]["host"], port=data["mysql"]["port"],
                            user=data["mysql"]["user"], password=data["mysql"]["password"])

# Consul
CONSUL_HOST = data["consul"]["host"]
CONSUL_PORT = data["consul"]["port"]

SERVICE_NAME = data["name"]
SERVICE_TAGS = data["tags"]
