# learning-python

Python学习笔记

目录说明

- [jupyter](jupyter)：使用jupyter notebook开发Python3
- [base](base)：[日志用法](base/log.py) [命令行参数](base/cmdline.py) [json](base/json-sample.py) [守护进程](base/daemon.py) [单元测试](base/base_test.py) [排序](base/sort2.py)
- [mysql](mysql)：使用MySQLdb访问MySQL数据库
- [redis](redis)：访问redis
- [mongodb](mongodb)：使用pymongo访问mongodb
- [restful-api](restful-api)：使用Flask构建Restful API
- [weibo](weibo): 调用新浪微博API

除非注明，否则均使用Python2.7

# 基础环境

macOS
```
brew install pycodestyle
alias pep8='_pycodestyle(){ /Users/pyw/Library/Python/2.7/bin/pycodestyle $1;}; _pycodestyle'
```

# 参考资料

- [pep8](https://www.python.org/dev/peps/pep-0008/)
- https://docs.python.org/2/library/json.html
- https://docs.python.org/zh-cn/3/library/getopt.html
- [《Dive Into Python》](https://book.douban.com/subject/1440658/)
- [3Blue1Brown](https://www.3blue1brown.com/) [Youtube频道](https://www.youtube.com/channel/UCYO_jab_esuFRV4b17AJtAw)