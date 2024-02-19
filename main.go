package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc-lab/db"
	"grpc-lab/memo"
	pb "grpc-lab/proto/gen/go"
	"log"
	"log/slog"
	"net"
	"net/http"
	"os"
	"time"
)

func startGRPCServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	memoServer := &memo.MemoServiceServer{}
	pb.RegisterMemoServiceServer(s, memoServer)

	log.Println("gRPC server listening at", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 요청이 들어온 시간
		start := time.Now()

		// 실제 핸들러 함수 실행
		next.ServeHTTP(w, r)

		// 요청 처리 후 로그 기록
		log.Printf("%s %s %s", r.Method, r.RequestURI, time.Since(start))
	})
}

func startGRPCGateway() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := pb.RegisterMemoServiceHandlerFromEndpoint(ctx, mux, "localhost:50051", opts); err != nil {
		log.Fatalf("Failed to register gRPC gateway: %v", err)
	}

	// 새로운 HTTP Mux 생성
	httpMux := http.NewServeMux()

	// Swagger UI와 JSON 파일에 대한 경로를 httpMux에 직접 등록
	fs := http.FileServer(http.Dir("./web/swagger-ui/dist"))
	httpMux.Handle("/swagger-ui/", http.StripPrefix("/swagger-ui/", fs))
	httpMux.HandleFunc("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./proto/gen/openapiv2/memo.swagger.json")
	})

	// gRPC-Gateway mux를 httpMux에 등록
	httpMux.Handle("/", mux)

	// 로깅 미들웨어 적용
	loggedMux := loggingMiddleware(httpMux)

	log.Println("gRPC-Gateway server listening at", ":8080")
	if err := http.ListenAndServe(":8080", loggedMux); err != nil {
		log.Fatalf("Failed to serve gRPC-Gateway: %v", err)
	}
}

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true})))

	// 데이터베이스 연결 초기화
	if err := db.InitDB("user:password@tcp(localhost:3306)/grpc_lab"); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// gRPC 서버 시작
	go startGRPCServer()

	// gRPC-Gateway 시작
	startGRPCGateway()
}
