import requests

headers = {
    "contentType": "application/json"
}


# 发送post请求
def register(name, id, address, port):
    url = "http://localhost:8500/v1/agent/service/register"
    rsp = requests.put(url, headers=headers, json={
        "Name": name,
        "ID": id,
        "Tags": ["shop-web", "bobby", "starry", "web"],
        "Address": address,
        "Port": port,
        "Check": {
            "GRPC": f"{address}:{port}",
            "GRPCUseTLS": False,
            "Timeout": "5s",
            "Interval": "5s",
            "DeregisterCriticalServiceAfter": "15s"
        }
    })
    if rsp.status_code == 200:
        print("注册成功")
    else:
        print(f"注册失败：{rsp.status_code}")


def de_register(id):
    url = f"http://localhost:8500/v1/agent/service/deregister/{id}"
    rsp = requests.put(url, headers=headers)
    if rsp.status_code == 200:
        print("注销成功")
    else:
        print(f"注销失败：{rsp.status_code}")


def filter_service(name):
    url = "http://localhost:8500/v1/agent/services"
    params = {
        "filter": f'Service == "{name}"'
    }
    rsp = requests.get(url, params=params).json()
    for key, value in rsp.items():
        print(key)


if __name__ == '__main__':
    register("shop-web", "shop-web", "localhost", 3000)
    de_register("shop-web")
    filter_service("shop-web")
