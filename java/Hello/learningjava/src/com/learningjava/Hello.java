package com.learningjava;

import javax.swing.text.NumberFormatter;
import java.awt.*;
import java.text.NumberFormat;
import java.util.Arrays;
import java.util.Date;

public class Hello {
    public static void var_base() {
        System.out.println("------------ var_base ------------");

        byte b = 1; // -128 ~ 127
        short s = 0; // -32K~32K
        int age = 35; // -2B~2B
        int bigint = 123_456_789;
        long bignum = 1_123_456_789L;
        float price = 18.99F;
        double price2 = 18.99;
        char letter = 'A';
        boolean isTrue = false;

        Date now = new Date();
        System.out.println(now);

        Point p1 = new Point(1, 1);
        Point p2 = p1;
        p1.x = 2;
        System.out.println(p2); // java.awt.Point[x=2,y=1]

        String msg1 = new String("Hello,\"AI\".");
        String msg2 = "Hello" + "!!";
        System.out.println(msg1); // Hello
        System.out.println(msg2.endsWith("!!")); // true
        System.out.println(msg2.startsWith("!!")); // false
        System.out.println(msg2.length()); // 7
        System.out.println(msg2.indexOf("H")); // 0
        System.out.println(msg2.indexOf("FF")); // -1
        System.out.println(msg2.replace("!", "~")); // Hello~~
        System.out.println(msg2); // Hello!!
        System.out.println(msg2.toLowerCase()); // hello!!

        final float PI = 3.14F;
        // PI = 1; // wrong, is readonly
        System.out.println(PI);

        // Implicit casting
        short ss = 1;
        int i = ss + 2;
        System.out.println(i); // 3
    }

    public static void conv() {
        double dx = 1.1;
        double dy = 2 + dx;
        System.out.println(dy); // 3.1

        String x = "1";
        int y = Integer.parseInt(x) + 1;
        System.out.println(y); // 2

        int result = Math.round(1.1F);
        System.out.println(result); // 1
        result = (int)Math.ceil(1.1F);
        System.out.println(result); // 2
        result = (int)Math.floor(1.1F);
        System.out.println(result); // 1
        result = (int)Math.max(1, 2);
        System.out.println(result); // 2
        double r = Math.random();
        System.out.println(r);
    }
    public static void array() {
        System.out.println("------------ array -------------");
        int[] users = new int[3];
        users[0] = 0;
        users[1] = 1;
        // users[2] default is 0;
        System.out.println(users);
        System.out.println(Arrays.toString(users));

        int[] users2 = {3, 1, 2};
        System.out.println(users2.length);
        Arrays.sort(users2); // change users2
        System.out.println(Arrays.toString(users2)); // [1, 2, 3]

        int[][] numbers  = new int[2][3]; // row=2;col=3;
        numbers[0][0] = 0;
        numbers[0][1] = 1;
        System.out.println(Arrays.deepToString(numbers)); // [[0, 1, 0], [0, 0, 0]]

        int[][] numbers2 = {{1,2,3},{4,5,6}};
        System.out.println(Arrays.deepToString(numbers2)); // [[1, 2, 3], [4, 5, 6]]
    }
    public static void var_operator() {
        int val = 10 / 3;
        System.out.println(val); // 3
        double dval = (double)10 / (double)3;
        System.out.println(dval); // 3.3333333333333335
        int x = 1;
        x++;
        ++x;
        x += 1;
        System.out.println(x); // 4
    }
    public static void fmt() {
        NumberFormat currency = NumberFormat.getCurrencyInstance();
        String ret = currency.format(1234567.891);
        System.out.println(ret); // ￥1,234,567.89
        NumberFormat percent = NumberFormat.getPercentInstance();
        ret = percent.format(0.1);
        System.out.println(ret); // 10%
    }
    public static void main(String[] args) {
        var_base();
        array();
        var_operator();
        conv();
        fmt();
    }
}
