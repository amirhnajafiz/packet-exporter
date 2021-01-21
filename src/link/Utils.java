package link;

import java.io.File;
import java.util.ArrayList;

public class Utils
{
    public static ArrayList<File> history;

    public static void utilsInit()
    {
        history = new ArrayList<>();
    }

    public static File open(String name)
    {
        for (File f : history)
        {
            if (f.getName().equals(name))
            {
                return f;
            }
        }
        return null;
    }

    public static void add(String path)
    {
        history.add(new File(path));
    }
}
