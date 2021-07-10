#!/usr/bin/env python3
#-*- coding: UTF-8 -*-
#coding=utf8

import tensorflow as tf

# random_normal返回形状为(1,4) 的张量。它的4个元素符合均值为100，标准差为0.35的正态分布
w = tf.Variable(initial_value=tf.random.normal(shape=(1,4), mean=100, stddev=0.35), name="w")
b = tf.Variable(tf.zeros([4]), name="b")

print (w)
print (b)

sess = tf.compat.v1.Session()
# 初始化为0
print (sess.run(tf.compat.v1.global_variables_initializer()))

# 0
print (sess.run([w, b]))

# 1
print (sess.run(tf.compat.v1.assign_add(b, [1, 1, 1, 1])))

# 验证b的值是否被加上了
print (sess.run(b))


# 保存
saver = tf.compat.v1.train.Saver({'w':w, 'b':b})
print (saver)

# global_step:训练到第几步
saver.save(sess, './save/1.save', global_step=0)

print (sess.run(tf.compat.v1.assign_add(b, [1, 1, 1, 1])))
print (sess.run(b))


saver.restore(sess, './save/1.save')
print (sess.run(b))

# 恢复数据流图
# tf.train.import_mete_graph











