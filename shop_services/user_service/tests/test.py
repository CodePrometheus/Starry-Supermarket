import grpc

from shop_services.user_service.proto import user_pb2_grpc, user_pb2


class UserTest:
    def __init__(self):
        # conn
        channel = grpc.insecure_channel("127.0.0.1:9000")
        self.stub = user_pb2_grpc.UserStub(channel)

    def user_list(self):
        rsp: user_pb2.UserListResponse = self.stub.GetUserList(user_pb2.PageInfo(
            pn=2, pSize=2
        ))
        print(rsp.total)
        for user in rsp.data:
            print(user.mobile, user.nickName)

    def create_user(self, nick_name, mobile, password):
        rsp: user_pb2.UserInfoResponse = self.stub.CreateUser(
            user_pb2.CreateUserInfo(nickName=nick_name,
                                    mobile=mobile,
                                    password=password))
        print(rsp.id)


if __name__ == '__main__':
    user = UserTest()
    user.create_user("starry", "12345678901", "admin")
