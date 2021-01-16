#!/usr/bin/env python
# -*- coding: UTF-8 -*-
# coding=utf8


import requests


class http_method():
    def get(self, url):
        r = requests.get(url, headers=None, params=None, data=None)
        r.status_code
        r.encoding
        r.content # bytes object
        # r.json() # if return value is json
        return r.text


if __name__ == "__main__":
    http = http_method()
    ret = http.get("https://www.github.com/")
    print(ret)
