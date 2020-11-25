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
curl "http://127.0.0.1:8081/api/v1.0/objects?token=ABCDEF0123456789"
```

获取第一个对象的详细信息（GET）
```
curl "http://127.0.0.1:8081/api/v1.0/objects/1?token=ABCDEF0123456789"
```

添加一个对象（POST）
```
curl -X POST -H 'Content-Type: application/json' -d '{"title":"title","description":"some text"}' 'http://127.0.0.1:8081/api/v1.0/objects?token=ABCDEF0123456789'
```

修改这个对象的信息（PUT）
```
curl -X PUT -H 'Content-Type: application/json' -d '{"title":"newtitle","description":"newdesc"}' 'http://127.0.0.1:8081/api/v1.0/objects/1?token=ABCDEF0123456789'
```

删除这个对象（DELETE）
```
curl -X DELETE "http://127.0.0.1:8081/api/v1.0/objects/1?token=ABCDEF0123456789"
```

## 说明

参数获取
- GET请求：通过flask.request.args获取；
- POST请求：通过flask.request.form获取；

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
