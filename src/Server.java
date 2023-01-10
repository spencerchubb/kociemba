package src;

import java.io.IOException;
import java.io.OutputStream;
import java.net.InetSocketAddress;

import com.sun.net.httpserver.HttpExchange;
import com.sun.net.httpserver.HttpHandler;
import com.sun.net.httpserver.HttpServer;

public class Server {

    public static void main(String[] args)throws Exception {
        int port = 4000;
        HttpServer server = HttpServer.create(new InetSocketAddress(port), 0);
        server.createContext("/solve", new SolveHandler());
        server.setExecutor(null);
        server.start();
        System.out.println("Server started on port " + port);
    }

    static class SolveHandler implements HttpHandler {
        @Override
        public void handle(HttpExchange t) throws IOException {
            // /solve/facelets
            String path = t.getRequestURI().getPath();

            // should be length 3
            // 0: ""
            // 1: "solve"
            // 2: facelets
            String[] pathSegments = path.split("/");

            if (pathSegments.length != 3) {
                t.sendResponseHeaders(400, 0);
                OutputStream os = t.getResponseBody();
                os.close();
                return;
            }

            String facelets = pathSegments[2];

            Search search = new Search();
            String solution = search.solution(facelets, 21, 500, 0, 0);

            t.sendResponseHeaders(200, solution.length());
            OutputStream os = t.getResponseBody();
            os.write(solution.getBytes());
            os.close();
        }
    }
}