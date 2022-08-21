#!/usr/bin/env python
# -*- coding: UTF-8 -*-
# coding=utf8

def test_readline():
    """readline包含回车"""
    f = open("input.py")
    line = f.readline()
    while line:
        print(line, len(line))
        break # only read one line
    f.close()
    last = line[-1:]
    print(last, ord(last)) # 10
    assert(last == '\n')

def test_input():
    content = input("input a string:")
    print(content, len(content)) # abc-->3

def main():
    test_input()
    test_readline()

if __name__ == '__main__':
    main()