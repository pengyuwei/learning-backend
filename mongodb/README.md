# 基础环境

```
$ python --version
Python 2.7.17

$ mongod --version
db version v3.6.3
git version: 9586e557d54ef70f9ca4b43c26892cd55257e1a5
OpenSSL version: OpenSSL 1.1.1  11 Sep 2018
allocator: tcmalloc
modules: none
build environment:
    distarch: x86_64
    target_arch: x86_64
    
pip install pymongo
```

验证

```
$ mongo
> show dbs
admin       0.000GB
config      0.000GB
db          0.000GB
local       0.000GB
testdb      0.000GB
> use testdb
switched to db testdb
> show collections
testcol
> db.testcol.find({tags:['2020','blog']})
{ "_id" : ObjectId("5fc2eb74c8498022b981dabb"), "date" : ISODate("2020-11-29T08:29:50.609Z"), "val" : { "A" : 1, "B" : "is b" }, "name" : "0xFF", "descr" : "Some text.", "tags" : [ "2020", "blog" ] }
```

# 参考文档

- https://zhuanlan.zhihu.com/p/51171906