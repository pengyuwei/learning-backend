package com.memcd.java;

public class BitTest {

    public static void main(String[] args) {
        // 10 10 --> 10 10 00
        System.out.println(10<<2); // 40

        // 10 10 --> 10
        System.out.println(10>>2); // 2

        // - 10 10 --> - 10 10 00
        System.out.println(-10<<2); // -40

        // - 10 10 --> - 10
        int a = -10;
        a = a >> 2;
        System.out.println(a); // -3

        // 无符号右移: 10 10 --> 00 10
        System.out.println(10 >>> 2); // 2

        // and: 10 10 ^ 00 10 --> 00 10
        System.out.println(10 & 2); // 2

        // or: 10 10 ^ 00 10 --> 10 10
        System.out.println(10 | 2); // 10

        // xor: 10 10 ^ 00 10 --> 10 00
        System.out.println(10 ^ 2); // 8

        // not
        System.out.println(~10); // -11
    }
}
