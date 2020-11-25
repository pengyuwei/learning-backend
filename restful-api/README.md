# Restful-API

参数获取
- GET请求：通过flask.request.args获取；
- POST请求：通过flask.request.form获取；

请求规范：
- 参数均为application/json

服务响应规范：
- 服务器回应的 HTTP 头的Content-Type属性统一设为application/json