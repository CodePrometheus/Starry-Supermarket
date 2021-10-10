import consul

c = consul.Consul(host="localhost")

address = "localhost"
port = 3000
check = {
    "GRPC": f"{address}:{port}",
    "GRPCUseTLS": False,
    "Timeout": "5s",
    "Interval": "5s",
    "DeregisterCriticalServiceAfter": "15s"
}

rsp = c.agent.services()
for key, val in rsp.items():
    rsp = c.agent.service.deregister(key)
print(rsp)
