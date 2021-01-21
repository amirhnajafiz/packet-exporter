package link;

import java.io.File;
import java.io.IOException;
import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.time.Duration;

public class Main
{
    public static void main(String[] args)
    {
        File output = new File("result.html");
        HttpClient client = HttpClient.newHttpClient();
        HttpRequest request = HttpRequest.newBuilder()
                .uri(URI.create("https://www.google.com"))
                .timeout(Duration.ofMinutes(1))
                .build();
        try
        {
            HttpResponse response = client.send(request, HttpResponse.BodyHandlers.ofFile(output.toPath()));
            System.out.println(response.statusCode() + " " + response.version());
        } catch (IOException | InterruptedException e)
        {
            e.printStackTrace();
        }
    }
}
