#! -*- coding:utf-8 -*-

import threading
import urllib2

def get_body(i):
	print "start",i
	urllib2.urlopen("http://cn.bing.com")
	print "end",i


for i in range(3):
	t = threading.Thread(target=get_body,args=(i,))
	t.start()