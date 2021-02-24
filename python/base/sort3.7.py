#!/nsr/bin/env python
# -*- coding: UTF-8 -*-
# coding=utf8

# sudo apt install python3.7
# python3.7 -m pip install numpy
# python3.7 -m pip install sortednp
import random
import time
import datetime
import numpy as np
import sortednp  # Merge and intersect sorted numpy arrays.

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
        a = next(it_a)
        b = next(it_b)
        while True:
            if a <= b:
                yield a; a = None
                a = next(it_a)
            else:
                yield b; y = None
                b = next(it_b)
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


# by Whistler 2021-2-24
def merge_numpy(a, b):
    c = sortednp.merge(a, b)
    return c
  

def nowns():
    return time.time_ns()
    # t = time.time.time_ns()
    # return int(round(t * 1000000))


def main():
    """
c_merge 1000 times spend 5086015 ns
  numpy 1000 times spend 13364420 ns
   sort 1000 times spend 30804664 ns
  merge 1000 times spend 281948268 ns
 naive2 1000 times spend 276838519 ns
   join 1000 times spend 243146704 ns
    """
    a = sorted(random.sample(range(10000), 1000))
    b = sorted(random.sample(range(10000), 1000))
  
    npa = np.array(a)
    npb = np.array(b)
    begin = nowns()
    for i in range(0, 1000):
        merge_numpy(npa, npb)
    end = nowns()
    print('  numpy 1000 times spend', end - begin, 'ns')

    begin = nowns()
    for i in range(0, 1000):
        sort(a, b)
    end = nowns()
    print('   sort 1000 times spend', end - begin, 'ns')

    begin = nowns()
    for i in range(0, 1000):
        merge(a, b)
    end = nowns()
    print('  merge 1000 times spend', end - begin, 'ns')

    begin = nowns()
    for i in range(0, 1000):
        naive2(a, b)
    end = nowns()
    print(' naive2 1000 times spend', end - begin, 'ns')

    begin = nowns()
    for i in range(0, 1000):
        join_sort(a, b)
    end = nowns()
    print('   join 1000 times spend', end - begin, 'ns')


if __name__ == '__main__':
    main()