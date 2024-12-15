package main

import (
	"log"
	"net"

	"github.com/jahidxuddin/uni-drop/internal/filetransfer"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000: %v", err)
	}

	grpcServer := grpc.NewServer()

	filetransfer.RegisterFileTransferServer(grpcServer, &filetransfer.Server{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000: %v", err)
	}
}
