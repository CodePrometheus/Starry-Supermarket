import json
import time

import nacos

SERVER_ADDRESSES = "localhost:8848"
NAMESPACE = "77a22bd7-4ce5-4af5-89f0-b427b6d6c55f"

client = nacos.NacosClient(SERVER_ADDRESSES, namespace=NAMESPACE,
                           username="nacos", password="nacos")

# get config
data_id = "user_svs"
group = "dev"
config = client.get_config(data_id, group)
print(type(config))
json_data = json.loads(config)
print(json_data)


def handler_cb(args):
    print("配置文件产生变化")
    print(args)


if __name__ == '__main__':
    client.add_config_watcher(data_id, group, handler_cb)
    time.sleep(3000)
