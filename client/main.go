package main

import (
	"context"
	"log"
	"time"

	"github.com/vietpham102301/grpc/pb" // Import package đã generate
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 1. Thiết lập kết nối đến Server gRPC
	// insecure.NewCredentials() vì chúng ta chưa dùng SSL/TLS
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Không thể kết nối: %v", err)
	}
	defer conn.Close()

	// 2. Tạo một Client từ kết nối trên
	c := pb.NewGreeterClient(conn)

	// 3. Chuẩn bị dữ liệu và Context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 4. Gọi hàm SayHello như một hàm Go bình thường
	res, err := c.SayHello(ctx, &pb.HelloRequest{Name: "Viet Pham"})
	if err != nil {
		log.Fatalf("Lỗi khi gọi hàm: %v", err)
	}

	// 5. In kết quả trả về từ Server
	log.Printf("Phản hồi từ Server: %s", res.GetMessage())
}
