import json

import consul
import grpc
from google.protobuf import empty_pb2

from shop_services.goods_service.proto import goods_pb2, goods_pb2_grpc
from shop_services.goods_service.settings import settings


class GoodsTest:
    def __init__(self):
        # 连接grpc服务器
        c = consul.Consul(host="localhost", port=8500)
        services = c.agent.services()
        ip = ""
        port = ""
        for key, value in services.items():
            if value["Service"] == settings.SERVICE_NAME:
                ip = value["Address"]
                port = value["Port"]
                break
        if not ip:
            raise Exception()
        channel = grpc.insecure_channel(f"{ip}:{port}")
        self.goods_stub = goods_pb2_grpc.GoodsStub(channel)

    def goods_list(self):
        rsp: goods_pb2.GoodsListResponse = self.goods_stub.GoodsList(
            goods_pb2.GoodsFilterRequest(keyWords="")
        )
        for item in rsp.data:
            print(item.name, item.shopPrice)

    def batch_get(self):
        ids = [421, 422]
        rsp: goods_pb2.GoodsListResponse = self.goods_stub.BatchGetGoods(
            goods_pb2.BatchGoodsIdInfo(id=ids)
        )
        for item in rsp.data:
            print(item.name, item.shopPrice)

    def get_detail(self, id):
        rsp = self.goods_stub.GetGoodsDetail(goods_pb2.GoodInfoRequest(
            id=id
        ))
        print(rsp.name)

    def category_list(self):
        rsp = self.goods_stub.GetAllCategoryList(empty_pb2.Empty())
        data = json.loads(rsp.jsonData)
        print(data)


if __name__ == "__main__":
    goods = GoodsTest()
    goods.category_list()
