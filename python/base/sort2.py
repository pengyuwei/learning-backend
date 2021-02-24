#!/usr/bin/env python
# -*- coding: UTF-8 -*-
# coding=utf8

import random
import time
import datetime


# by zzj 2021-2-24
def naive2(a, b):
    assert a and b
    i = j = k = 0
    len_a, len_b = len(a), len(b)
    result = [None] * (len_a + len_b)
    while True:
        if i == len_a: break
        if j == len_b: break
        if a[i] <= b[j]:
            result[k] = a[i]
            k += 1; i += 1
        else:
            result[k] = b[j]
            k += 1; j += 1
    # append the remainder
    for x in a[i:] + b[j:]:
        result[k] = x
        k += 1
    return result


# by Dieken 2021-2-24
def merge(a, b):
    la = len(a)
    lb = len(b)
    c = [None] * (la + lb)
    i = 0
    j = 0
    k = 0
    while i < la and j < lb:
        if a[i] < b[j]:
            c[k] = a[i]
            i = i + 1
        else:
            c[k] = b[j]
            j = j + 1
        k = k + 1
    return c


# by zzj 2021-2-24
def join(a, b):
    """takes two sorted list and yields a sorted sequence"""
    it_a = iter(a)
    it_b = iter(b)
    a = b = None
    try:
        a = it_a.next()
        b = it_b.next()
        while True:
            if a <= b:
                yield a; a = None
                a = it_a.next()
            else:
                yield b; y = None
                b = it_b.next()
    except StopIteration:
        if a or b: yield a or b
        for a in it_a: yield a
        for b in it_b: yield b


def join_sort(a, b):
    result = []
    for i in join(a, b):
        result.append(i)
    return result


def sort(a, b):
    return sorted(a + b)


def nowus():
    t = time.time()
    return int(round(t * 1000000))


def main():
    """
 numpy 1000 times spend 13148 us
  sort 1000 times spend 50459 us
 merge 1000 times spend 213063 us
naive2 1000 times spend 259805 us
  join 1000 times spend 355014 us
    """
    a = sorted(random.sample(range(10000), 1000))
    b = sorted(random.sample(range(10000), 1000))
  
    begin = nowus()
    for i in range(0, 1000):
        sort(a, b)
    end = nowus()
    print '  sort 1000 times spend', end - begin, 'us'

    begin = nowus()
    for i in range(0, 1000):
        merge(a, b)
    end = nowus()
    print ' merge 1000 times spend', end - begin, 'us'

    begin = nowus()
    for i in range(0, 1000):
        naive2(a, b)
    end = nowus()
    print 'naive2 1000 times spend', end - begin, 'us'

    begin = nowus()
    for i in range(0, 1000):
        join_sort(a, b)
    end = nowus()
    print '  join 1000 times spend', end - begin, 'us'

if __name__ == '__main__':
    main()