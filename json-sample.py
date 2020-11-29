#!/usr/bin/env python
# -*- coding: UTF-8 -*-
# coding=utf8

import json


def test_json():
    j = {"count": 1, "Nums":[1,2,3]}
    s = json.dumps(j, encoding="utf-8")
    s2 = json.dumps(j, indent=4)
    print(s)
    print(s2)

    s = '{"db":{"host":"localhost","port":"3306","user":"test","password":"test","db":"test"}}'
    j = json.loads(s, encoding="utf-8")
    print(j)

    fd = open('tmp.json', 'w')
    fd.write(s2)
    fd.close()

    with open('tmp.json') as json_file:
        config = json.load(json_file)
        print(config)

if __name__ == '__main__':
    test_json()