#!/usr/bin/env python
# -*- coding: UTF-8 -*-
# coding=utf8


import getopt
import sys


def usage():
    print("Usage:")
    print("./cmdline.py -a -b -c10 -d hello p1 p2")
    print("./cmdline.py --help")
    print("-" * 40)


def main():
    print("-" * 40)
    try:
        opts, args = getopt.getopt(sys.argv[1:], "ho:v", ["help", "output="])
    except getopt.GetoptError as err:
        print(err)
        usage()
        sys.exit(2)
    output = None
    verbose = False
    for o, a in opts:
        if o == "-v":
            verbose = True
        elif o in ("-h", "--help"):
            usage()
            sys.exit()
        elif o in ("-o", "--output"):
            output = a
        else:
            assert False, "unhandled option"


if __name__ == '__main__':
    try:
        opts, args = getopt.getopt(sys.argv[1:], "abc:d:")
        usage()
        for op, value in opts:
            print(op, value)
        print "args is ", args
        sys.exit(0)
    except getopt.GetoptError as err:
        pass
    main()
