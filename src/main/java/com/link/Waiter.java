package com.link;

import java.util.Scanner;

public class Waiter
{
    private static Scanner scanner;

    public static void waiterInit()
    {
        scanner = new Scanner(System.in);
    }

    public static String getOrderByString()
    {
        return scanner.next();
    }
}
