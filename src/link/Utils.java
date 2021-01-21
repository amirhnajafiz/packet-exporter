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

    public static File add(String path)
    {
        File file = new File(path);
        history.add(file);
        return file;
    }
}
