package test;

import java.net.http.HttpClient;
import java.net.http.HttpResponse;
import java.net.http.HttpRequest;

import java.time.Duration;

/*
 * The server must be running before TestServer will pass.
 */
public class TestServer {

    public static void main(String[] args) throws Exception {
        String facelets = "DUUBULDBFRBFRRULLLBRDFFFBLURDBFDFDRFRULBLUFDURRBLBDUDL";
        String expected = "R2 U2 B2 L2 F2 U' L2 R2 B2 R2 D B2 F L' F U2 F' R' D' L2 R'";

        HttpClient client = HttpClient.newHttpClient();
        HttpRequest request = HttpRequest.newBuilder()
                .GET()
                .uri(java.net.URI.create("http://localhost:4000/solve/" + facelets))
                .timeout(Duration.ofSeconds(2))
                .build();
        HttpResponse<String> response = client.send(request, HttpResponse.BodyHandlers.ofString());
        String solution = response.body();

        if (!solution.equals(expected)) {
            System.out.println("Got " + solution + ", expected " + expected);
            return;
        }
        System.out.println("Test passed");
    }
}
