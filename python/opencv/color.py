#!/usr/bin/env python3
# -*- coding: UTF-8 -*-
# coding=utf8

import cv2

"""
pip3 install opencv-python
"""

def main():
    img = cv2.imread('demo.png')
    cv2.imshow('demo.png',img)
    cv2.waitKey(0)
    cv2.imwrite('new.png', img)

if __name__ == '__main__':
    main()