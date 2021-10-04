from peewee import *

from shop_services.user_service.settings import db


class BaseModel(Model):
    class Meta:
        database = db.DB


class User(BaseModel):
    # 用户模型
    GENDER_CHOICES = (
        ("female", "女"),
        ("male", "男")
    )

    ROLE_CHOICES = (
        (1, "普通用户"),
        (2, "管理员")
    )

    # 手机号码定为用户的唯一标识
    mobile = CharField(max_length=11, index=True, unique=True, verbose_name="手机号码")

    password = CharField(max_length=100, verbose_name="密码")  # 1. 密文 2. 密文不可反解
    nick_name = CharField(max_length=20, null=True, verbose_name="昵称")
    head_url = CharField(max_length=200, null=True, verbose_name="头像")
    birthday = DateField(null=True, verbose_name="生日")
    address = CharField(max_length=200, null=True, verbose_name="地址")
    desc = TextField(null=True, verbose_name="个人简介")
    gender = CharField(max_length=6, choices=GENDER_CHOICES, null=True, verbose_name="性别")
    role = IntegerField(default=1, choices=ROLE_CHOICES, verbose_name="用户角色")


if __name__ == '__main__':
    # db.DB.create_tables([User])
    # m = hashlib.md5()
    # m.update(b"12345")
    # print(m.hexdigest())
    from passlib.handlers.pbkdf2 import pbkdf2_sha256

    for i in range(10):
        user = User()
        user.nick_name = f"starry-{i}"
        user.mobile = f"1387001234{i}"
        user.password = pbkdf2_sha256.hash("starry")
        user.save()
