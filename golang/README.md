# Go语言学习笔记

每个文件对是一个知识点的实验

初始环境：
```
export GO111MODULE=on
export GOPROXY=https://goproxy.io
go mod init learning-go
go get github.com/jinzhu/gorm
go get github.com/go-sql-driver/mysql
go get github.com/gorilla/mux
go get github.com/mongodb/mongo-go-driver
```

运行
```
go run xxxx.go
```

# 参考文档

语言基础
- 《Go语言入门经典》
- https://www.runoob.com/go/go-tutorial.html
- http://c.biancheng.net/view/61.html
- https://www.runoob.com/go/go-structures.html

MySQL
- https://gorm.io/zh_CN/docs/index.html

Redis
- https://github.com/go-redis/redis
- https://www.runoob.com/redis/redis-tutorial.html

Context
- https://godoc.org/context
- http://c.biancheng.net/view/5714.html

Mongo
- [mongo/quickstart](https://www.mongodb.com/blog/search/golang%20quickstart)
- [mongo-driver](https://godoc.org/go.mongodb.org/mongo-driver/mongo)
- [mongo-driver/bson](https://godoc.org/go.mongodb.org/mongo-driver/bson)
- [Quick Start: Golang & MongoDB - Modeling Documents with Go Data Structures](https://www.mongodb.com/blog/post/quick-start-golang--mongodb--modeling-documents-with-go-data-structures)
- [mongo-go-driver详细使用示例](https://www.cnblogs.com/zcqkk/p/11234227.html)

Moudle && pkg
- [Go模块代理速度实测](https://studygolang.com/topics/9994)
- [Go手动安装pkg包](https://blog.csdn.net/u012393450/article/details/82665588)