
## Setting up the Environment

To set up your development environment for Grpc and Grpc-Gateway, ensure you have the correct versions of the protobuf compiler and the necessary plugins installed. Run the following commands to install the required tools:

```shell
$ go mod tidy

$ go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc \
    github.com/fullstorydev/grpcui/cmd/grpcui@latest
```

These tools will help you compile your `.proto` files into Go code, generate Grpc-Gateway reverse-proxy code, and serve your Grpc services with a Swagger UI and grpcui for easy testing and interaction.

## Running the HTTP Server

After setting up your environment and generating the necessary code, you can run the HTTP server to start serving your Grpc and HTTP requests. The server setup involves starting both the Grpc server and the Grpc-Gateway, optionally with Swagger UI and grpcui support for API documentation and testing.

### Starting the Grpc Server

The Grpc server is started by the `startGrpcServer` function, which sets up the server on a specified port, registers the service implementations, and starts listening for incoming Grpc requests.

### Starting the Grpc-Gateway

The Grpc-Gateway acts as a reverse-proxy, translating RESTful HTTP APIs into Grpc calls. It's started by the `startGrpcGateway` function, which also configures the HTTP multiplexer to serve Swagger UI and grpcui if enabled.

To start the HTTP server with Swagger UI and grpcui enabled, use the following command:

```go
startGrpcGateway(true, true)
```

This command initializes the Grpc-Gateway, sets up the Swagger UI at `/swagger-ui/`, and enables grpcui at `/grpc-ui/`, allowing for easy testing and interaction with your Grpc services through a web interface.

### Running the Server

To run your server, execute the `main` function in your server's main package. This function initializes the database connection, starts the Grpc server in a separate goroutine, and starts the Grpc-Gateway. Ensure that the database and any other external dependencies are correctly set up and accessible.

For more detailed information on configuring and running the server, refer to the source code documentation and comments.
