package main

import (
	"context"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/vietpham102301/grpc/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Chào " + in.GetName() + "! Bạn đang gọi qua gRPC hoặc REST đó."}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Không thể listen port 50051: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	log.Println("gRPC Server đang chạy tại port :50051...")
	go func() {
		log.Fatal(s.Serve(lis))
	}()

	// --- BẮT ĐẦU CHẠY GRPC-GATEWAY (REST PROXY) ---
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:50051",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Không thể kết nối tới gRPC server:", err)
	}

	gwmux := runtime.NewServeMux()
	// Đăng ký Greeter handler
	err = pb.RegisterGreeterHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Không thể đăng ký Gateway:", err)
	}

	log.Println("REST Gateway đang chạy tại port :8080...")
	log.Fatal(http.ListenAndServe(":8080", gwmux))
}
