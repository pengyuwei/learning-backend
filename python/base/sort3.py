#!/usr/bin/env python
# -*- coding: UTF-8 -*-
# coding=utf8

# python3
# pip3 install sortednp
# python3 sort3.py;python sort2.py
import random
import time
import numpy as np
import sortednp


# by Whistler 2021-2-24
def merge_numpy(a, b):
    c = sortednp.merge(a, b)
    return c


def nowus():
    t = time.time()
    return int(round(t * 1000000))


if __name__ == '__main__':
    """
 numpy 1000 times spend 13336 us
  sort 1000 times spend 51874 us
 merge 1000 times spend 214143 us
naive2 1000 times spend 250681 us
  join 1000 times spend 313955 us
    """
    ra = sorted(random.sample(range(10000), 1000))
    rb = sorted(random.sample(range(10000), 1000))
    a = np.array(ra)
    b = np.array(rb)

    begin = nowus()
    for i in range(1, 1000):
        merge_numpy(a, b)
    end = nowus()
    print(' numpy 1000 times spend', end - begin, 'us')
