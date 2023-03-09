#!/usr/bin/env python
# -*- coding: UTF-8 -*-
# coding=utf8

"""
鸭子类型: 如果走起路来像鸭子, 叫起来也像鸭子, 那么它就是鸭子
鸭子类型是动态类型语言中的一种设计风格，一个对象的特征不是由父类决定，而是通过对象的方法决定的。
"""

class Duck:
    def call(self):
        print("GuaGua!")

class Bird:
    def call(self):
        print("JiuJiu~")

class Doge:
    def call(self):
        print("WangWangWang!")

def solo(duck):
    duck.call()

duck = Duck()
bird = Bird()
doge = Doge()

for pet in [duck, bird, doge]:
    solo(pet)
