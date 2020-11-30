# Restful-API

## 开发环境

```
python --version
Python 2.7.16

pip install flask
pip install flask-cors
pip install pyOpenSSL
pip install Flask-SSLify
```

启动服务
```
python api.py
```

## 本地测试

登录（POST）

```
curl -d "user=test&passwd=test" http://127.0.0.1:8081/api/v1.0/get_token
```

获取全部对象列表（GET）
```
curl -H 'Content-Type: application/json' -H 'Authorization:ABCDEF0123456789' "http://127.0.0.1:8081/api/v1.0/objects"
```

获取第一个对象的详细信息（GET）
```
curl -H 'Content-Type: application/json' -H 'Authorization:ABCDEF0123456789' "http://127.0.0.1:8081/api/v1.0/objects/1"
```

添加一个对象（POST）
```
curl -X POST -H 'Content-Type: application/json' -H 'Authorization:ABCDEF0123456789' -d '{"title":"title","description":"some text"}' 'http://127.0.0.1:8081/api/v1.0/objects'
```

```
客户端尝试
OPTIONS /api/v1.0/objects HTTP/1.1
Host: example.com:8081
Connection: keep-alive
Accept: */*
Access-Control-Request-Method: POST
Access-Control-Request-Headers: authorization,content-type
Origin: null
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.67 Safari/537.36
Sec-Fetch-Mode: cors
Accept-Encoding: gzip, deflate
Accept-Language: zh-CN,zh;q=0.9

服务器响应
HTTP/1.0 200 OK
Content-Type: text/html; charset=utf-8
Allow: POST, HEAD, OPTIONS, GET
Access-Control-Allow-Origin: null
Access-Control-Allow-Methods: DELETE, GET, HEAD, OPTIONS, PATCH, POST, PUT
Vary: Origin
Access-Control-Allow-Headers: authorization, content-type
Access-Control-Allow-Credentials: true
Content-Length: 0
Server: Werkzeug/1.0.1 Python/2.7.16
Date: Mon, 30 Nov 2020 14:07:01 GMT

客户端再次请求
POST /api/v1.0/objects HTTP/1.1
Host: example.com:8081
Connection: keep-alive
Content-Length: 41
User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.67 Safari/537.36
DNT: 1
authorization: ABCDEF0123456789
content-type: application/json
Accept: */*
Origin: null
Accept-Encoding: gzip, deflate
Accept-Language: zh-CN,zh;q=0.9

title=newtitle&description=newdescription
```

修改这个对象的信息（PUT）
```
curl -X PUT -H 'Content-Type: application/json' -H 'Authorization:ABCDEF0123456789' -d '{"title":"newtitle","description":"newdesc"}' 'http://127.0.0.1:8081/api/v1.0/objects/1'
```

删除这个对象（DELETE）
```
curl -X DELETE -H 'Content-Type: application/json' -H 'Authorization:ABCDEF0123456789' "http://127.0.0.1:8081/api/v1.0/objects/1"
```

## 说明

参数获取
- URL参数请求：通过flask.request.args获取；
- FORM表单请求：通过flask.request.form获取；
- Header头：通过flask.request.headers获取；

请求规范：
- 参数均为application/json

服务响应规范：
- 服务器回应的 HTTP 头的Content-Type属性统一设为application/json

如果使用https，这样启动服务：
```
app.run(debug=True, host="0.0.0.0", port=80, ssl_context=(
    "server/server-cert.pem",
    "server/server-key.pem")
)
```

# 参考文档

- https://developer.mozilla.org/zh-CN/docs/Web/API/Fetch_API/Using_Fetch#Headers
- https://itbilu.com/javascript/js/VkiXuUcC.html
- https://hyjk2000.github.io/2015/04/02/cors-for-restful-api/
- https://www.ruanyifeng.com/blog/2016/04/cors.html