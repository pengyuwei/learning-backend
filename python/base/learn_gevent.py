#! -*- coding:utf-8 -*-

import urllib2
import gevent
from gevent import monkey


def get_body(i):
	print "start",i
	urllib2.urlopen("http://cn.bing.com")
	print "end",i

monkey.patch_all()
tasks = [gevent.spawn(get_body,i) for i in range(3)]
gevent.joinall(tasks)