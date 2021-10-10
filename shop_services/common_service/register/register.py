import random

import consul
import requests

from shop_services.common_service.register import base


class ConsulRegister(base.Register):
    def __init__(self, host, port):
        self.host = host
        self.port = port
        self.consul = consul.Consul(host, port)

    def register(self, name, id, addr, port, tags, check) -> bool:
        if check is None:
            check = {
                "GRPC": f"{addr}:{port}",
                "GRPCUseTLS": False,
                "Timeout": "5s",
                "Interval": "5s",
                "DeregisterCriticalServiceAfter": "15s"
            }
        else:
            check = check

        return self.consul.agent.service.register(name=name, service_id=id,
                                                  address=addr, port=port, tags=tags, check=check)

    def deregister(self, service_id):
        return self.consul.agent.service.deregister(service_id)

    def get_all_service(self):
        return self.consul.agent.services()

    def filter_service(self, filter):
        url = f"http://{self.host}:{self.port}/v1/agent/services"
        params = {
            "filter": filter
        }
        return requests.get(url, params=params).json()

    def get_host_port(self, filter):
        url = f"http://{self.host}:{self.port}/v1/agent/services"
        params = {
            "filter": filter
        }
        data = requests.get(url, params=params).json()
        if data:
            service_info = random.choice(list(data.values()))
            return service_info["Address"], service_info["Port"]
        return None, None
