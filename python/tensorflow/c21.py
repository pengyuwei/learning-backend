#!/usr/bin/env python3
#-*- coding: UTF-8 -*-
#coding=utf8

# C21-22 操作Operation
# https://github.com/DjangoPeng/tensorflow-101/blob/master/notebook-examples/chapter-3/operation.ipynb
# https://github.com/tensorflow/tensorflow/tree/r1.14/tensorflow/python/ops
import tensorflow as tf


# 常量操作
a = tf.constant(2)
b = tf.constant(3)

# 创建会话，并执行计算操作
with tf.compat.v1.Session() as sess:
    print("a: %i" % sess.run(a))
    print("b: %i" % sess.run(b))
    print("Addition with constants: %i" % sess.run(a + b))
    print("Multiplication with constants: %i" % sess.run(a * b))

# 占位符操作
# tf.reset_default_graph()
# TF V1 code
#x = tf.compat.v1.placeholder(tf.int32, shape=(1024, 1024))
x = tf.compat.v1.placeholder(dtype=tf.int32, shape=(), name="x")
y = tf.compat.v1.placeholder(dtype=tf.int32, shape=(), name="y")
# TF V2 code
#x = tf.Variable(tf.zeros(shape=()), name="x")
#y = tf.Variable(tf.zeros(shape=()), name="y")

# 计算操作
add = tf.add(x, y)
mul = tf.multiply(x, y)

# 加载默认数据流图
with tf.compat.v1.Session() as sess:
    # 填充数据后，执行操作
    print("Addition with variables: %d" % sess.run(add, feed_dict={x: 10, y: 5}))
    print("Multiplication with variables: %d" % sess.run(mul, feed_dict={x: 2, y: 3}))
    
