# 基础环境

1. 微博开放平台注册账号，创建网站应用：
https://open.weibo.com/

配置文件：temp.json:
```
{
    "APP_KEY": "",
    "APP_SECRET": "",
    "CALLBACK_URL": "",
}
```

2. 部署基础环境
```
git clone --depth=1 https://github.com/michaelliao/sinaweibopy.git
sudo python3 setup.py install
$ python --version
Python 3.6.9
```

# 参考文档

- https://open.weibo.com/wiki/%E5%BE%AE%E5%8D%9AAPI
- http://michaelliao.github.io/sinaweibopy/
- https://github.com/michaelliao/sinaweibopy