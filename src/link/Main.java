package link;

import java.io.File;
import java.util.Scanner;

public class Main
{
    public static void messageInConsole(String message)
    {
        System.out.print(message);
    }

    public static void main(String[] args)
    {
        Utils.utilsInit();
        Downloader downloader = new Downloader();
        Scanner scanner = new Scanner(System.in);
        messageInConsole("Enter url > ");
        String link = scanner.next();
        messageInConsole("Enter output > ");
        String output = scanner.next();
        File output_file = Utils.open(output);
        if (output_file == null)
        {
            output_file = Utils.add(output);
        }
        boolean result = downloader.touchLink(link, output_file);
        if (result)
        {
            messageInConsole("OK");
        } else
        {
            messageInConsole("FAILED");
        }
    }
}
