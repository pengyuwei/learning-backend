#!/usr/bin/env python
# -*- coding: UTF-8 -*-
# coding=utf8

"""
Python协程示例
"""


def display():
    r = ''
    while True:
        n = yield r  # 获取send的参数n，并且返回r

        if not n:
            return
        print('[----]display %d' % n)
        r = 'Next'
    print('You can not see me.')


def sendMessage(c):
    c.send(None)  # 启动生成器
    n = 0
    while n < 5:
        n = n + 1
        print('[SEND]send %d' % n)
        r = c.send(n)
        print('[SEND]get: %s' % r)
    c.close()


def demo1():
    d = display  # 和协程用法无关，仅为对比
    print(d)  # <function display at 0x1065b38c0>

    c = display()  # 因为函数里包含yield，因此这里解释器并不会掉用（也并不是获取函数指针）
    print(c)  # <generator object display at 0x10e64acd0>
    sendMessage(c)


if __name__ == '__main__':
    demo1()
