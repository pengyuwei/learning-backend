#!/usr/bin/env python
# -*- coding: UTF-8 -*-
# coding=utf8

import redis
import time


def test_redis():
    r = redis.Redis(host='localhost', port=6379, decode_responses=True)

    r.set('name', '0xFF')
    print(r['name'])
    print(r.get('name'))
    print(type(r.get('name')))

    # Overdue key
    key = 'overdue-key'
    r.set(key, 'value = timeout 3s', ex=3)
    print(r.get(key))
    r.expire(key, 2)
    time.sleep(5)
    print(r.get(key))


if __name__ == '__main__':
    test_redis()
