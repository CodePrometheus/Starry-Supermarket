import datetime
import logging

from peewee import *

logger = logging.getLogger("peewee")
logger.setLevel(logging.DEBUG)
logger.addHandler(logging.StreamHandler())

db = MySQLDatabase('starry-supermarket', host='localhost',
                   user='root', passwd='root')


class BaseModel(Model):
    add_time = DateTimeField(default=datetime.datetime.now, verbose_name="添加时间")

    class Meta:
        database = db  # 这里是数据库链接，为了方便建立多个表，可以把这个部分提炼出来形成一个新的类


class Person(BaseModel):
    first = CharField()
    last = CharField()

    class Meta:
        primary_key = CompositeKey('first', 'last')


class Pet(BaseModel):
    owner_first = CharField()
    owner_last = CharField()
    pet_name = CharField()

    class Meta:
        constraints = [SQL('FOREIGN KEY(owner_first, owner_last) REFERENCES person(first, last)')]


class Blog(BaseModel):
    pass


class Tag(BaseModel):
    pass


# 复合主键
class BlogToTag(BaseModel):
    """A simple "through" table for many-to-many relationship."""
    blog = ForeignKeyField(Blog)
    tag = ForeignKeyField(Tag)

    class Meta:
        primary_key = CompositeKey('blog', 'tag')


class User(BaseModel):
    # 如果没有设置主键，那么自动生成一个id的主键
    username = CharField(max_length=20)
    age = CharField(default=18, max_length=20, verbose_name="年龄")

    class Meta:
        table_name = 'new_user'  # 这里可以自定义表名


if __name__ == "__main__":
    db.connect()
    db.create_tables([Person, Pet, Blog, Tag, BlogToTag, User])

    # id = Person.insert({
    #     'first': 'star',
    #     'last': '2021'
    # }).execute()
    #
    # for i in range(10):
    #     User.create(username=f"2021-{i}", age=random.randint(18, 40))
    #
    # id = Blog.insert({}).execute()
    # print(id)

    # and
    person = Person.select().where((Person.first == "li") & (Person.first == "li1"))
    print(person.sql())
