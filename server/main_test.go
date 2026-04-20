package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/vietpham102301/grpc/pb"
)

func TestSayHello(t *testing.T) {
	s := &server{}

	req := &pb.HelloRequest{
		Name: "Viet Pham",
	}
	res, err := s.SayHello(context.Background(), req)

	assert.NoError(t, err)
	expectedMessage := "Chào Viet Pham! Bạn đang gọi qua gRPC hoặc REST đó."
	assert.Equal(t, expectedMessage, res.GetMessage())
}
