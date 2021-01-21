package link;

import java.io.File;

public class Main
{
    public static void messageInConsole(String message)
    {
        System.out.print(message);
    }

    public static void main(String[] args)
    {
        Utils.utilsInit();
        Waiter.waiterInit();

        Downloader downloader = new Downloader();
        messageInConsole("Enter url > ");
        String link = Waiter.getOrderByString();

        messageInConsole("Enter output > ");
        String output = Waiter.getOrderByString();
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
