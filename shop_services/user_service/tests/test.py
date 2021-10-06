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
        for i in rsp.data:
            print(i.mobile, i.nickname)

    def create_user(self, nick_name, mobile, password):
        rsp: user_pb2.UserInfoResponse = self.stub.CreateUser(
            user_pb2.CreateUserInfo(nickname=nick_name,
                                    mobile=mobile,
                                    password=password))
        print(rsp.id)

    def get_user_by_id(self, request_id):
        rsp: user_pb2.UserInfoResponse = self.stub.GetUserById(
            user_pb2.IdRequest(id=request_id)
        )
        print(rsp.mobile)

    def update_user(self, request_id, nickname, gender, birthday):
        rsp: user_pb2.UserInfoResponse = self.stub.UpdateUser(
            user_pb2.UpdateUserInfo(id=request_id,
                                    nickname=nickname,
                                    gender=gender,
                                    birthday=birthday)
        )
        print(rsp)


if __name__ == '__main__':
    user = UserTest()
    user.user_list()
    # user.create_user("starry", "12345678901", "admin")
    # user.get_user_by_id(19)
    # user.update_user(19, nickname="admin", gender="1", birthday=1641319300)
