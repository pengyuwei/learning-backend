# learning-python

Python学习笔记

目录说明

- [jupyter](jupyter)：使用jupyter notebook开发Python3
- [base](base)：[日志用法](base/log.py) [命令行参数](base/cmdline.py) [json](base/json-sample.py) [守护进程](base/daemon.py) [单元测试](base/base_test.py) [排序](base/sort2.py)
- [thread](thread): 进程、线程、[协程](thread/coroutines.py)
- [mysql](mysql)：使用MySQLdb访问MySQL数据库
- [redis](redis)：访问redis
- [mongodb](mongodb)：使用pymongo访问mongodb
- [restful-api](restful-api)：使用Flask构建Restful API
- [weibo](weibo): 调用新浪微博API
- [abnormal](abnormal): 异常数据检测

除非注明，否则均使用Python2.7

## Environment

macOS
```
brew install pycodestyle
alias pep8='_pycodestyle(){ /Users/pyw/Library/Python/2.7/bin/pycodestyle $1;}; _pycodestyle'
```

## Virtual environments

```
sudo apt install python3.10-venv
mkdir -p venv
cd venv
python3 -m venv .
cd bin
source activate
pip install torch --extra-index-url https://download.pytorch.org/whl/cu113
deactivate
```

## 常用

```
__bases__：以元组返回一个类所直接继承的类。
__mor__：  以元组返回继承关系链。
__class__：返回对象所属的类。
__globals__：以字典返回函数所在模块命名空间中的所有变量。
__subclasses__()：以列表的返回类的子类。
__builtins__：内建函数
```

## Reference

- [pep8](https://www.python.org/dev/peps/pep-0008/)
- https://docs.python.org/2/library/json.html
- https://docs.python.org/zh-cn/3/library/getopt.html
- [《Dive Into Python》](https://book.douban.com/subject/1440658/)
- [3Blue1Brown](https://www.3blue1brown.com/) [Youtube频道](https://www.youtube.com/channel/UCYO_jab_esuFRV4b17AJtAw)