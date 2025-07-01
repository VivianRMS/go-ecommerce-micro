package account.java_client;

import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;
import io.grpc.StatusRuntimeException;
import account.v1.AccountServiceGrpc.AccountServiceBlockingStub;
import account.v1.AccountServiceGrpc;
import account.v1.AccountOuterClass.PostAccountRequest;
import account.v1.AccountOuterClass.PostAccountResponse;

public class JavaClient {
    public static void main(String[] args) {
        System.out.println("[Java gRPC Client] Starting request...");

        // Build channel to server
        ManagedChannel channel = ManagedChannelBuilder.forAddress("localhost", 8080)
                .usePlaintext()  // disables TLS
                .build();

        // Create stub
        AccountServiceBlockingStub stub = AccountServiceGrpc.newBlockingStub(channel);

        // Build request
        PostAccountRequest request = PostAccountRequest.newBuilder()
                .setName("Java")
                .build();

        // Call RPC
        try {
            PostAccountResponse response = stub.postAccount(request);
            System.out.println("Received response: " + response);
        } catch (StatusRuntimeException e) {
            System.err.println("gRPC Error: " + e.getStatus());
        } finally {
            channel.shutdown();
        }
    }
}
