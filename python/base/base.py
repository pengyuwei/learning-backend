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

    # @mysql_conn
    # def func3(self, conn):
    #     assert(conn == 'mysql')


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


def test_fstring():
    """>=Python3.6"""
    name = "0xFF"
    address = "github.com"
    print(f"我是 {name}, 来自 {address}")
    print(f"1 + 1 = {1 + 1}")  # 1 + 1 = 2
    print(f"sum([1, 2, 3]) = {sum([1, 2, 3])}") # sum([1, 2, 3]) = 6
    a = lambda x: x + 100
    print(f"{a}") # <function>
    print(f"{a(1)}") # 101

def test_print():
    var = "10"
    print("Here are " + var + " test cases.")
    print("Here are %s test cases." % var)
    print("Here are %d%d test cases." % (1, 0))
    print("Here are {0} test cases.".format(var))

def main():
    test_language_base()
    test_dict()
    test_staticmethod()
    test_fstring()
    test_print()

    demo = CDemo()
    # demo.func3()


if __name__ == '__main__':
    main()
