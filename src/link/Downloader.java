package link;

import java.io.File;
import java.io.IOException;
import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.nio.file.Path;
import java.time.Duration;

public class Downloader
{
    public boolean touchLink(String link, File output)
    {
        HttpClient client = HttpClient.newHttpClient();
        HttpRequest request = HttpRequest.newBuilder()
                .uri(URI.create(link))
                .timeout(Duration.ofMinutes(1))
                .build();
        try
        {
            HttpResponse<Path> response = client.send(request, HttpResponse.BodyHandlers.ofFile(output.toPath()));
            System.out.println(response.statusCode() + " " + response.version());
            System.out.println(response.headers());
            return true;
        } catch (IOException | InterruptedException e)
        {
            e.printStackTrace();
            return false;
        }
    }
}
