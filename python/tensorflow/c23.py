#!/usr/bin/env python3
#-*- coding: UTF-8 -*-
#coding=utf8

# C23 Session
# https://github.com/DjangoPeng/tensorflow-101/blob/master/notebook-examples/chapter-3/operation.ipynb
# https://github.com/tensorflow/tensorflow/tree/r1.14/tensorflow/python
import tensorflow as tf


x = tf.compat.v1.placeholder(dtype=tf.float32)
w = tf.compat.v1.Variable(1.0)
b = tf.compat.v1.Variable(1.0)
y = w * x + b

with tf.compat.v1.Session() as sess:
    tf.global_variables_initializer().run()
    fetch = y.eval(feed_dict={x: 3.0}) # 估算张量
    print(fetch)
    
