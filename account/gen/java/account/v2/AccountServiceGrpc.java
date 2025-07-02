package account.v2;

import static io.grpc.MethodDescriptor.generateFullMethodName;

/**
 */
@javax.annotation.Generated(
    value = "by gRPC proto compiler (version 1.58.0)",
    comments = "Source: account/v2/account.proto")
@io.grpc.stub.annotations.GrpcGenerated
public final class AccountServiceGrpc {

  private AccountServiceGrpc() {}

  public static final java.lang.String SERVICE_NAME = "account.v2.AccountService";

  // Static method descriptors that strictly reflect the proto.
  private static volatile io.grpc.MethodDescriptor<account.v2.AccountOuterClass.PostAccountRequest,
      account.v2.AccountOuterClass.PostAccountResponse> getPostAccountMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "PostAccount",
      requestType = account.v2.AccountOuterClass.PostAccountRequest.class,
      responseType = account.v2.AccountOuterClass.PostAccountResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<account.v2.AccountOuterClass.PostAccountRequest,
      account.v2.AccountOuterClass.PostAccountResponse> getPostAccountMethod() {
    io.grpc.MethodDescriptor<account.v2.AccountOuterClass.PostAccountRequest, account.v2.AccountOuterClass.PostAccountResponse> getPostAccountMethod;
    if ((getPostAccountMethod = AccountServiceGrpc.getPostAccountMethod) == null) {
      synchronized (AccountServiceGrpc.class) {
        if ((getPostAccountMethod = AccountServiceGrpc.getPostAccountMethod) == null) {
          AccountServiceGrpc.getPostAccountMethod = getPostAccountMethod =
              io.grpc.MethodDescriptor.<account.v2.AccountOuterClass.PostAccountRequest, account.v2.AccountOuterClass.PostAccountResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "PostAccount"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  account.v2.AccountOuterClass.PostAccountRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  account.v2.AccountOuterClass.PostAccountResponse.getDefaultInstance()))
              .setSchemaDescriptor(new AccountServiceMethodDescriptorSupplier("PostAccount"))
              .build();
        }
      }
    }
    return getPostAccountMethod;
  }

  private static volatile io.grpc.MethodDescriptor<account.v2.AccountOuterClass.GetAccountRequest,
      account.v2.AccountOuterClass.GetAccountResponse> getGetAccountMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetAccount",
      requestType = account.v2.AccountOuterClass.GetAccountRequest.class,
      responseType = account.v2.AccountOuterClass.GetAccountResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<account.v2.AccountOuterClass.GetAccountRequest,
      account.v2.AccountOuterClass.GetAccountResponse> getGetAccountMethod() {
    io.grpc.MethodDescriptor<account.v2.AccountOuterClass.GetAccountRequest, account.v2.AccountOuterClass.GetAccountResponse> getGetAccountMethod;
    if ((getGetAccountMethod = AccountServiceGrpc.getGetAccountMethod) == null) {
      synchronized (AccountServiceGrpc.class) {
        if ((getGetAccountMethod = AccountServiceGrpc.getGetAccountMethod) == null) {
          AccountServiceGrpc.getGetAccountMethod = getGetAccountMethod =
              io.grpc.MethodDescriptor.<account.v2.AccountOuterClass.GetAccountRequest, account.v2.AccountOuterClass.GetAccountResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetAccount"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  account.v2.AccountOuterClass.GetAccountRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  account.v2.AccountOuterClass.GetAccountResponse.getDefaultInstance()))
              .setSchemaDescriptor(new AccountServiceMethodDescriptorSupplier("GetAccount"))
              .build();
        }
      }
    }
    return getGetAccountMethod;
  }

  private static volatile io.grpc.MethodDescriptor<account.v2.AccountOuterClass.GetAccountsRequest,
      account.v2.AccountOuterClass.GetAccountsResponse> getGetAccountsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "GetAccounts",
      requestType = account.v2.AccountOuterClass.GetAccountsRequest.class,
      responseType = account.v2.AccountOuterClass.GetAccountsResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<account.v2.AccountOuterClass.GetAccountsRequest,
      account.v2.AccountOuterClass.GetAccountsResponse> getGetAccountsMethod() {
    io.grpc.MethodDescriptor<account.v2.AccountOuterClass.GetAccountsRequest, account.v2.AccountOuterClass.GetAccountsResponse> getGetAccountsMethod;
    if ((getGetAccountsMethod = AccountServiceGrpc.getGetAccountsMethod) == null) {
      synchronized (AccountServiceGrpc.class) {
        if ((getGetAccountsMethod = AccountServiceGrpc.getGetAccountsMethod) == null) {
          AccountServiceGrpc.getGetAccountsMethod = getGetAccountsMethod =
              io.grpc.MethodDescriptor.<account.v2.AccountOuterClass.GetAccountsRequest, account.v2.AccountOuterClass.GetAccountsResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "GetAccounts"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  account.v2.AccountOuterClass.GetAccountsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  account.v2.AccountOuterClass.GetAccountsResponse.getDefaultInstance()))
              .setSchemaDescriptor(new AccountServiceMethodDescriptorSupplier("GetAccounts"))
              .build();
        }
      }
    }
    return getGetAccountsMethod;
  }

  private static volatile io.grpc.MethodDescriptor<account.v2.AccountOuterClass.ListAccountsRequest,
      account.v2.AccountOuterClass.ListAccountsResponse> getListAccountsMethod;

  @io.grpc.stub.annotations.RpcMethod(
      fullMethodName = SERVICE_NAME + '/' + "ListAccounts",
      requestType = account.v2.AccountOuterClass.ListAccountsRequest.class,
      responseType = account.v2.AccountOuterClass.ListAccountsResponse.class,
      methodType = io.grpc.MethodDescriptor.MethodType.UNARY)
  public static io.grpc.MethodDescriptor<account.v2.AccountOuterClass.ListAccountsRequest,
      account.v2.AccountOuterClass.ListAccountsResponse> getListAccountsMethod() {
    io.grpc.MethodDescriptor<account.v2.AccountOuterClass.ListAccountsRequest, account.v2.AccountOuterClass.ListAccountsResponse> getListAccountsMethod;
    if ((getListAccountsMethod = AccountServiceGrpc.getListAccountsMethod) == null) {
      synchronized (AccountServiceGrpc.class) {
        if ((getListAccountsMethod = AccountServiceGrpc.getListAccountsMethod) == null) {
          AccountServiceGrpc.getListAccountsMethod = getListAccountsMethod =
              io.grpc.MethodDescriptor.<account.v2.AccountOuterClass.ListAccountsRequest, account.v2.AccountOuterClass.ListAccountsResponse>newBuilder()
              .setType(io.grpc.MethodDescriptor.MethodType.UNARY)
              .setFullMethodName(generateFullMethodName(SERVICE_NAME, "ListAccounts"))
              .setSampledToLocalTracing(true)
              .setRequestMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  account.v2.AccountOuterClass.ListAccountsRequest.getDefaultInstance()))
              .setResponseMarshaller(io.grpc.protobuf.ProtoUtils.marshaller(
                  account.v2.AccountOuterClass.ListAccountsResponse.getDefaultInstance()))
              .setSchemaDescriptor(new AccountServiceMethodDescriptorSupplier("ListAccounts"))
              .build();
        }
      }
    }
    return getListAccountsMethod;
  }

  /**
   * Creates a new async stub that supports all call types for the service
   */
  public static AccountServiceStub newStub(io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<AccountServiceStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<AccountServiceStub>() {
        @java.lang.Override
        public AccountServiceStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new AccountServiceStub(channel, callOptions);
        }
      };
    return AccountServiceStub.newStub(factory, channel);
  }

  /**
   * Creates a new blocking-style stub that supports unary and streaming output calls on the service
   */
  public static AccountServiceBlockingStub newBlockingStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<AccountServiceBlockingStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<AccountServiceBlockingStub>() {
        @java.lang.Override
        public AccountServiceBlockingStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new AccountServiceBlockingStub(channel, callOptions);
        }
      };
    return AccountServiceBlockingStub.newStub(factory, channel);
  }

  /**
   * Creates a new ListenableFuture-style stub that supports unary calls on the service
   */
  public static AccountServiceFutureStub newFutureStub(
      io.grpc.Channel channel) {
    io.grpc.stub.AbstractStub.StubFactory<AccountServiceFutureStub> factory =
      new io.grpc.stub.AbstractStub.StubFactory<AccountServiceFutureStub>() {
        @java.lang.Override
        public AccountServiceFutureStub newStub(io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
          return new AccountServiceFutureStub(channel, callOptions);
        }
      };
    return AccountServiceFutureStub.newStub(factory, channel);
  }

  /**
   */
  public interface AsyncService {

    /**
     */
    default void postAccount(account.v2.AccountOuterClass.PostAccountRequest request,
        io.grpc.stub.StreamObserver<account.v2.AccountOuterClass.PostAccountResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getPostAccountMethod(), responseObserver);
    }

    /**
     */
    default void getAccount(account.v2.AccountOuterClass.GetAccountRequest request,
        io.grpc.stub.StreamObserver<account.v2.AccountOuterClass.GetAccountResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetAccountMethod(), responseObserver);
    }

    /**
     */
    default void getAccounts(account.v2.AccountOuterClass.GetAccountsRequest request,
        io.grpc.stub.StreamObserver<account.v2.AccountOuterClass.GetAccountsResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getGetAccountsMethod(), responseObserver);
    }

    /**
     */
    default void listAccounts(account.v2.AccountOuterClass.ListAccountsRequest request,
        io.grpc.stub.StreamObserver<account.v2.AccountOuterClass.ListAccountsResponse> responseObserver) {
      io.grpc.stub.ServerCalls.asyncUnimplementedUnaryCall(getListAccountsMethod(), responseObserver);
    }
  }

  /**
   * Base class for the server implementation of the service AccountService.
   */
  public static abstract class AccountServiceImplBase
      implements io.grpc.BindableService, AsyncService {

    @java.lang.Override public final io.grpc.ServerServiceDefinition bindService() {
      return AccountServiceGrpc.bindService(this);
    }
  }

  /**
   * A stub to allow clients to do asynchronous rpc calls to service AccountService.
   */
  public static final class AccountServiceStub
      extends io.grpc.stub.AbstractAsyncStub<AccountServiceStub> {
    private AccountServiceStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected AccountServiceStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new AccountServiceStub(channel, callOptions);
    }

    /**
     */
    public void postAccount(account.v2.AccountOuterClass.PostAccountRequest request,
        io.grpc.stub.StreamObserver<account.v2.AccountOuterClass.PostAccountResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getPostAccountMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getAccount(account.v2.AccountOuterClass.GetAccountRequest request,
        io.grpc.stub.StreamObserver<account.v2.AccountOuterClass.GetAccountResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetAccountMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void getAccounts(account.v2.AccountOuterClass.GetAccountsRequest request,
        io.grpc.stub.StreamObserver<account.v2.AccountOuterClass.GetAccountsResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getGetAccountsMethod(), getCallOptions()), request, responseObserver);
    }

    /**
     */
    public void listAccounts(account.v2.AccountOuterClass.ListAccountsRequest request,
        io.grpc.stub.StreamObserver<account.v2.AccountOuterClass.ListAccountsResponse> responseObserver) {
      io.grpc.stub.ClientCalls.asyncUnaryCall(
          getChannel().newCall(getListAccountsMethod(), getCallOptions()), request, responseObserver);
    }
  }

  /**
   * A stub to allow clients to do synchronous rpc calls to service AccountService.
   */
  public static final class AccountServiceBlockingStub
      extends io.grpc.stub.AbstractBlockingStub<AccountServiceBlockingStub> {
    private AccountServiceBlockingStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected AccountServiceBlockingStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new AccountServiceBlockingStub(channel, callOptions);
    }

    /**
     */
    public account.v2.AccountOuterClass.PostAccountResponse postAccount(account.v2.AccountOuterClass.PostAccountRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getPostAccountMethod(), getCallOptions(), request);
    }

    /**
     */
    public account.v2.AccountOuterClass.GetAccountResponse getAccount(account.v2.AccountOuterClass.GetAccountRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetAccountMethod(), getCallOptions(), request);
    }

    /**
     */
    public account.v2.AccountOuterClass.GetAccountsResponse getAccounts(account.v2.AccountOuterClass.GetAccountsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getGetAccountsMethod(), getCallOptions(), request);
    }

    /**
     */
    public account.v2.AccountOuterClass.ListAccountsResponse listAccounts(account.v2.AccountOuterClass.ListAccountsRequest request) {
      return io.grpc.stub.ClientCalls.blockingUnaryCall(
          getChannel(), getListAccountsMethod(), getCallOptions(), request);
    }
  }

  /**
   * A stub to allow clients to do ListenableFuture-style rpc calls to service AccountService.
   */
  public static final class AccountServiceFutureStub
      extends io.grpc.stub.AbstractFutureStub<AccountServiceFutureStub> {
    private AccountServiceFutureStub(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      super(channel, callOptions);
    }

    @java.lang.Override
    protected AccountServiceFutureStub build(
        io.grpc.Channel channel, io.grpc.CallOptions callOptions) {
      return new AccountServiceFutureStub(channel, callOptions);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<account.v2.AccountOuterClass.PostAccountResponse> postAccount(
        account.v2.AccountOuterClass.PostAccountRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getPostAccountMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<account.v2.AccountOuterClass.GetAccountResponse> getAccount(
        account.v2.AccountOuterClass.GetAccountRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetAccountMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<account.v2.AccountOuterClass.GetAccountsResponse> getAccounts(
        account.v2.AccountOuterClass.GetAccountsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getGetAccountsMethod(), getCallOptions()), request);
    }

    /**
     */
    public com.google.common.util.concurrent.ListenableFuture<account.v2.AccountOuterClass.ListAccountsResponse> listAccounts(
        account.v2.AccountOuterClass.ListAccountsRequest request) {
      return io.grpc.stub.ClientCalls.futureUnaryCall(
          getChannel().newCall(getListAccountsMethod(), getCallOptions()), request);
    }
  }

  private static final int METHODID_POST_ACCOUNT = 0;
  private static final int METHODID_GET_ACCOUNT = 1;
  private static final int METHODID_GET_ACCOUNTS = 2;
  private static final int METHODID_LIST_ACCOUNTS = 3;

  private static final class MethodHandlers<Req, Resp> implements
      io.grpc.stub.ServerCalls.UnaryMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ServerStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.ClientStreamingMethod<Req, Resp>,
      io.grpc.stub.ServerCalls.BidiStreamingMethod<Req, Resp> {
    private final AsyncService serviceImpl;
    private final int methodId;

    MethodHandlers(AsyncService serviceImpl, int methodId) {
      this.serviceImpl = serviceImpl;
      this.methodId = methodId;
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public void invoke(Req request, io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        case METHODID_POST_ACCOUNT:
          serviceImpl.postAccount((account.v2.AccountOuterClass.PostAccountRequest) request,
              (io.grpc.stub.StreamObserver<account.v2.AccountOuterClass.PostAccountResponse>) responseObserver);
          break;
        case METHODID_GET_ACCOUNT:
          serviceImpl.getAccount((account.v2.AccountOuterClass.GetAccountRequest) request,
              (io.grpc.stub.StreamObserver<account.v2.AccountOuterClass.GetAccountResponse>) responseObserver);
          break;
        case METHODID_GET_ACCOUNTS:
          serviceImpl.getAccounts((account.v2.AccountOuterClass.GetAccountsRequest) request,
              (io.grpc.stub.StreamObserver<account.v2.AccountOuterClass.GetAccountsResponse>) responseObserver);
          break;
        case METHODID_LIST_ACCOUNTS:
          serviceImpl.listAccounts((account.v2.AccountOuterClass.ListAccountsRequest) request,
              (io.grpc.stub.StreamObserver<account.v2.AccountOuterClass.ListAccountsResponse>) responseObserver);
          break;
        default:
          throw new AssertionError();
      }
    }

    @java.lang.Override
    @java.lang.SuppressWarnings("unchecked")
    public io.grpc.stub.StreamObserver<Req> invoke(
        io.grpc.stub.StreamObserver<Resp> responseObserver) {
      switch (methodId) {
        default:
          throw new AssertionError();
      }
    }
  }

  public static final io.grpc.ServerServiceDefinition bindService(AsyncService service) {
    return io.grpc.ServerServiceDefinition.builder(getServiceDescriptor())
        .addMethod(
          getPostAccountMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              account.v2.AccountOuterClass.PostAccountRequest,
              account.v2.AccountOuterClass.PostAccountResponse>(
                service, METHODID_POST_ACCOUNT)))
        .addMethod(
          getGetAccountMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              account.v2.AccountOuterClass.GetAccountRequest,
              account.v2.AccountOuterClass.GetAccountResponse>(
                service, METHODID_GET_ACCOUNT)))
        .addMethod(
          getGetAccountsMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              account.v2.AccountOuterClass.GetAccountsRequest,
              account.v2.AccountOuterClass.GetAccountsResponse>(
                service, METHODID_GET_ACCOUNTS)))
        .addMethod(
          getListAccountsMethod(),
          io.grpc.stub.ServerCalls.asyncUnaryCall(
            new MethodHandlers<
              account.v2.AccountOuterClass.ListAccountsRequest,
              account.v2.AccountOuterClass.ListAccountsResponse>(
                service, METHODID_LIST_ACCOUNTS)))
        .build();
  }

  private static abstract class AccountServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoFileDescriptorSupplier, io.grpc.protobuf.ProtoServiceDescriptorSupplier {
    AccountServiceBaseDescriptorSupplier() {}

    @java.lang.Override
    public com.google.protobuf.Descriptors.FileDescriptor getFileDescriptor() {
      return account.v2.AccountOuterClass.getDescriptor();
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.ServiceDescriptor getServiceDescriptor() {
      return getFileDescriptor().findServiceByName("AccountService");
    }
  }

  private static final class AccountServiceFileDescriptorSupplier
      extends AccountServiceBaseDescriptorSupplier {
    AccountServiceFileDescriptorSupplier() {}
  }

  private static final class AccountServiceMethodDescriptorSupplier
      extends AccountServiceBaseDescriptorSupplier
      implements io.grpc.protobuf.ProtoMethodDescriptorSupplier {
    private final java.lang.String methodName;

    AccountServiceMethodDescriptorSupplier(java.lang.String methodName) {
      this.methodName = methodName;
    }

    @java.lang.Override
    public com.google.protobuf.Descriptors.MethodDescriptor getMethodDescriptor() {
      return getServiceDescriptor().findMethodByName(methodName);
    }
  }

  private static volatile io.grpc.ServiceDescriptor serviceDescriptor;

  public static io.grpc.ServiceDescriptor getServiceDescriptor() {
    io.grpc.ServiceDescriptor result = serviceDescriptor;
    if (result == null) {
      synchronized (AccountServiceGrpc.class) {
        result = serviceDescriptor;
        if (result == null) {
          serviceDescriptor = result = io.grpc.ServiceDescriptor.newBuilder(SERVICE_NAME)
              .setSchemaDescriptor(new AccountServiceFileDescriptorSupplier())
              .addMethod(getPostAccountMethod())
              .addMethod(getGetAccountMethod())
              .addMethod(getGetAccountsMethod())
              .addMethod(getListAccountsMethod())
              .build();
        }
      }
    }
    return result;
  }
}
