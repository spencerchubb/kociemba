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
        server.createContext("/test", new MyHandler());
        server.createContext("/solve", new SolveHandler());
        server.setExecutor(null); // creates a default executor
        server.start();
        System.out.println("Server started on port " + port);
    }

    static class MyHandler implements HttpHandler {
        @Override
        public void handle(HttpExchange t) throws IOException {
            String response = "This is the response";
            t.sendResponseHeaders(200, response.length());
            OutputStream os = t.getResponseBody();
            os.write(response.getBytes());
            os.close();
        }
    }

    static class SolveHandler implements HttpHandler {
        @Override
        public void handle(HttpExchange t) throws IOException {
            String query = t.getRequestURI().getQuery();
            String[] params = query.split("&");
            String[] pair = params[0].split("=");
            String key = pair[0];
            String value = pair[1];

            System.out.println(value);

            Search search = new Search();
            String solution = search.solution(value, 21, 500, 0, 0);
            System.out.println(solution);

            t.sendResponseHeaders(200, solution.length());
            OutputStream os = t.getResponseBody();
            os.write(solution.getBytes());
            os.close();
        }
    }
}