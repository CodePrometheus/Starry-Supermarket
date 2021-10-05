import time
from datetime import date

import grpc
from google.protobuf import empty_pb2
from loguru import logger
from passlib.hash import pbkdf2_sha256
from peewee import DoesNotExist

from shop_services.user_service.model.models import User
from shop_services.user_service.proto import user_pb2_grpc, user_pb2


# 定义的方法
class UserServicer(user_pb2_grpc.UserServicer):
    @staticmethod
    @logger.catch
    def convert_user_to_rsp(user):
        # 将user的model对象转换成rpc的message对象
        user_info_rsp = user_pb2.UserInfoResponse()

        user_info_rsp.id = user.id
        user_info_rsp.password = user.password
        user_info_rsp.mobile = user.mobile
        user_info_rsp.role = user.role

        if user.nick_name:
            user_info_rsp.nickname = user.nick_name
        if user.gender:
            user_info_rsp.gender = user.gender
        if user.birthday:
            user_info_rsp.birthday = int(time.mktime(user.birthday.timetuple()))

        return user_info_rsp

    @logger.catch
    def GetUserList(self, request: user_pb2.PageInfo, context):
        # 获取用户的列表
        rsp = user_pb2.UserListResponse()

        users = User.select()
        rsp.total = users.count()
        logger.info("用户列表Got")
        start = 0
        per_page_nums = 10

        # 分页
        if request.pSize:
            per_page_nums = request.pSize
        if request.pn:
            start = per_page_nums * (request.pn - 1)

        users = users.limit(per_page_nums).offset(start)

        for user in users:
            rsp.data.append(self.convert_user_to_rsp(user))

        return rsp

    @logger.catch
    def GetUserById(self, request: user_pb2.IdRequest, context):
        # 通过id查询用户
        try:
            user = User.get(User.id == request.id)
            return self.convert_user_to_rsp(user)
        except DoesNotExist as e:
            logger.error(e)
            context.set_code(grpc.StatusCode.NOT_FOUND)
            context.set_details("用户不存在")
            return user_pb2.UserInfoResponse()

    @logger.catch
    def GetUserByMobile(self, request: user_pb2.MobileRequest, context):
        # 通过mobile查询用户
        try:
            user = User.get(User.mobile == request.mobile)
            return self.convert_user_to_rsp(user)
        except DoesNotExist as e:
            logger.error(e)
            context.set_code(grpc.StatusCode.NOT_FOUND)
            context.set_details("用户不存在")
            return user_pb2.UserInfoResponse()

    @logger.catch
    def CreateUser(self, request: user_pb2.CreateUserInfo, context):
        # 新建用户
        try:
            User.get(User.mobile == request.mobile)
            context.set_code(grpc.StatusCode.ALREADY_EXISTS)
            context.set_details("用户已存在")
            return user_pb2.UserInfoResponse()
        except DoesNotExist as e:
            logger.error(e)
            pass

        user = User()
        user.nick_name = request.nickname
        user.mobile = request.mobile
        user.password = pbkdf2_sha256.hash(request.password)
        user.save()

        return self.convert_user_to_rsp(user)

    @logger.catch
    def UpdateUser(self, request: user_pb2.UpdateUserInfo, context):
        # 更新用户
        try:
            user = User.get(User.id == request.id)
            user.nick_name = request.nickname
            user.gender = request.gender
            user.birthday = date.fromtimestamp(request.birthday)
            print(request.birthday)
            print(user.birthday)
            user.save()
            return empty_pb2.Empty()
        except DoesNotExist:
            context.set_code(grpc.StatusCode.NOT_FOUND)
            context.set_details("用户不存在")
            return user_pb2.UserInfoResponse()

    @logger.catch
    def CheckPassword(self, request: user_pb2.PasswordCheckInfo):
        return user_pb2.CheckResponse(success=pbkdf2_sha256.verify(request.password, request.encryptedPassword))
