package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/fullstorydev/grpcui/standalone"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	"grpc-lab/db"
	"grpc-lab/memo"
	pb "grpc-lab/proto/gen/go"
)

var grpcPort = 50051

// startGrpcServer initializes and starts the gRPC server.
func startGrpcServer() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	memoServer := &memo.MemoServiceServer{}
	pb.RegisterMemoServiceServer(s, memoServer)

	// Register reflection service for grpc-ui.
	reflection.Register(s)

	log.Println("gRPC server listening at", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

// loggingMiddleware logs the HTTP requests.
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.RequestURI, time.Since(start))
	})
}

// startGrpcGateway starts the gRPC-Gateway, optionally enabling Swagger UI and grpc-ui.
func startGrpcGateway(enableSwaggerUI, enableGrpcUI bool) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	grpcServerEndpoint := fmt.Sprintf("127.0.0.1:%d", grpcPort)
	gatewayMux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := pb.RegisterMemoServiceHandlerFromEndpoint(ctx, gatewayMux, grpcServerEndpoint, opts); err != nil {
		log.Fatalf("Failed to register gRPC gateway: %v", err)
	}

	httpMux := http.NewServeMux()

	// Optionally serve Swagger UI.
	if enableSwaggerUI {
		fs := http.FileServer(http.Dir("./web/swagger-ui/dist"))
		httpMux.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui/", fs))
		httpMux.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "./proto/gen/openapiv2/memo.swagger.json")
		})
	}

	// Optionally serve grpc-ui.
	if enableGrpcUI {
		conn, err := grpc.Dial(grpcServerEndpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("Failed to dial server: %v", err)
		}
		defer conn.Close()

		grpcuiHandler, err := standalone.HandlerViaReflection(ctx, conn, grpcServerEndpoint)
		if err != nil {
			log.Fatalf("Failed to create grpcui handler: %v", err)
		}

		httpMux.Handle("/grpc-ui/", http.StripPrefix("/grpc-ui", grpcuiHandler))
	}

	// Serve the gRPC-Gateway.
	httpMux.Handle("/", gatewayMux)

	// Apply logging middleware.
	loggedMux := loggingMiddleware(httpMux)

	log.Println("gRPC-Gateway server listening at", ":8080")
	if err := http.ListenAndServe(":8080", loggedMux); err != nil {
		log.Fatalf("Failed to serve gRPC-Gateway: %v", err)
	}
}

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true})))

	// Initialize database connection.
	if err := db.InitDB("user:password@tcp(localhost:3306)/grpc_lab"); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Start the gRPC server in a new goroutine.
	go startGrpcServer()

	// Start the gRPC-Gateway with Swagger UI and grpc-ui enabled.
	startGrpcGateway(true, true)
}
