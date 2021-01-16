#!/usr/bin/env python
# -*- coding: UTF-8 -*-
# coding=utf8

import logging


def init_log():
    logging.basicConfig(
            level=logging.INFO,
            format='%(asctime)s %(filename)s'
                   '[line:%(lineno)d] %(levelname)s %(message)s',
            datefmt='%a, %d %b %Y %H:%M:%S',
            filename='log.log',
            filemode='w')
    console = logging.StreamHandler()
    console.setLevel(logging.DEBUG)
    console.setLevel(logging.DEBUG)
    formatter = logging.Formatter('%(name)-12s: %(levelname)-8s %(message)s')
    console.setFormatter(formatter)
    logging.getLogger('').addHandler(console)


if __name__ == '__main__':
    init_log()
    logging.info("A info log.")
    logging.debug("A debug log.")
