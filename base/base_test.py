#!/usr/bin/env python
# -*- coding: UTF-8 -*-
# coding=utf8

import unittest


class TestDict(unittest.TestCase):

    def setUp(self):
        print('setUp...')

    def tearDown(self):
        print('tearDown...')

    def test_func1(self):
        print('test_func1')
        self.assertEqual(1, 1)
        return True

    def test_func2(self):
        self.assertTrue(True)
        print('test_func2')
        return True

    def will_not_invoke(self):
        print("will_not_invoke")
        d = Dict()
        with self.assertRaises(AttributeError):
            value = d.empty


# python base_test.py
# python -m unittest base_test
if __name__ == '__main__':
    unittest.main()
