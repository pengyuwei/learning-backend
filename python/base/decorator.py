#!/usr/bin/env python
# -*- coding: UTF-8 -*-
# coding=utf8

def log(func):
    def wrapper(*args, **kw):
        print('------ %s() -------' % func.__name__)
        return func(True, *args, **kw)
    return wrapper

@log
def test_decorator(check, para):
    from datetime import datetime as DT
    print(DT.now())
    assert(check == True)
    assert(para == 1)

def mysql_conn(func):
    def wrapper_keep_conn(*args, **kwargs):
        conn = "mysql"
        ret = func(conn, *args, **kwargs)
        conn = ''
    return wrapper_keep_conn

@mysql_conn
def test_decorator2(conn, p1, p2):
    assert(conn == 'mysql')
    assert(p1 == 1)
    assert(p2 == 2)

def check_authorization(f):
    def wrapper(*args):
        print args[0].url
        token = args[0].url[-4:] 
        assert(token == '1234')
        return f(*args, token=token)
    return wrapper

class Client(object):
    def __init__(self, url):
        self.url = url

    @check_authorization
    def get(self, token):
        print 'get', token

def main():
    test_decorator(1)
    test_decorator2(1, 2)
    Client('http://www.google.com?token=1234').get()

if __name__ == '__main__':
    main()