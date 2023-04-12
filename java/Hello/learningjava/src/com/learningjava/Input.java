package com.learningjava;

import java.util.Scanner;

public class Input {
    public static void case1() {
        Scanner scanner = new Scanner(System.in);
        System.out.print("Age:");
        byte age = scanner.nextByte(); // max is 127
        System.out.println("Age: " + age);

        System.out.print("Name:");
        String name = scanner.next().trim(); // in="Steve Jobs", name="Steve"
        System.out.println("Name: " + name); // Steve
        name = scanner.next();
        System.out.println("Name: " + name); // Jobs

        System.out.print("Name:");
        name = scanner.nextLine(); // in="Steve Jobs", name="Steve Jobs"
        System.out.println("Name: " + name); // Steve Jobs
    }
    public static void case2() {
        String[] fruits = {"apple", "orange", "mango"};
        for (String fruit : fruits) {
            System.out.println(fruit);
        }
    }
    public static void main(String[] args) {
        case1();
        case2();
    }
}
