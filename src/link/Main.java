package link;

import java.util.Scanner;

public class Main
{
    public static void messageInConsole(String message)
    {
        System.out.print(message);
    }

    public static void main(String[] args)
    {
        Downloader downloader = new Downloader();
        Scanner scanner = new Scanner(System.in);
        messageInConsole("Enter url > ");
        String link = scanner.next();
        boolean result = downloader.touchLink(link);
        if (result)
        {
            messageInConsole("OK");
        } else
        {
            messageInConsole("FAILED");
        }
    }
}
