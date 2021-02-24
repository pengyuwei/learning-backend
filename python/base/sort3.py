#!/usr/bin/env python
# -*- coding: UTF-8 -*-
# coding=utf8

# python3
# pip3 install sortednp
# python3 sort3.py;python sort2.py
import random
import time
import numpy as np
import sortednp  # Merge and intersect sorted numpy arrays.


# by Whistler 2021-2-24
def merge_numpy(a, b):
    c = sortednp.merge(a, b)
    return c


def nowus():
    t = time.time()
    return int(round(t * 1000000))


if __name__ == '__main__':
    """
 numpy 1000 times spend 13477 us
  sort 1000 times spend 53067 us
 merge 1000 times spend 231770 us
naive2 1000 times spend 242620 us
  join 1000 times spend 302756 us
    """
    ra = sorted(random.sample(range(10000), 1000))
    rb = sorted(random.sample(range(10000), 1000))
    a = np.array(ra)
    b = np.array(rb)

    begin = nowus()
    for i in range(0, 1000):
        merge_numpy(a, b)
    end = nowus()
    print(' numpy 1000 times spend', end - begin, 'us')
