#!/usr/bin/env python
# -*- coding: UTF-8 -*-
# coding=utf8


class CDemo():
    val = 0

    def __init__(self, val=1):
        self.val = val

    @staticmethod
    def func():
        print('staticmethod\n')

    def func2(self):
        print('val=%d' % self.val)


def test_staticmethod():
    CDemo.func()

    obj = CDemo(2)
    obj.func()
    obj.func2()


def test_language_base():
    print(test_language_base.__name__)
    assert(test_language_base.__name__ == 'test_language_base')

    a, b = 1, 2
    c = a if a > b else b
    assert(c == b)
    print(dir(a))


def test_dict():
    a = [1, 2, 3, 4]
    b = a[::-1]
    print(a, b)
    assert(b[0] == a[3])

    list_str = "[1, 2, 3]"
    tmp = eval(list_str)
    assert(tmp[0] == 1)

    d = {i: i * i for i in range(10)}
    s = {i * 2 for i in range(10)}
    print(d, s)

    i = 0
    for i, item in enumerate(a):
        print(i, item)
    for i, item in enumerate(a, 1):
        print(i, item)
    print('-' * 40)


def log(func):
    def wrapper(*args, **kw):
        print('------ %s() -------' % func.__name__)
        return func(*args, **kw)
    return wrapper


@log
def test_decorator():
    from datetime import datetime as DT
    print(DT.now())


def main():
    test_language_base()
    test_dict()
    test_staticmethod()
    test_decorator()


if __name__ == '__main__':
    main()
