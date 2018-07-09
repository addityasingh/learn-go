package main

import (
	"fmt"
	"net"

	"github.com/Sirupsen/logrus"
	pb "github.com/addityasingh/protobuf/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type EchoServer struct {
}

func (s *EchoServer) Reply(ctx context.Context, request *pb.EchoRequest) (*pb.EchoResponse, error) {
	return &pb.EchoResponse{
		Message: request.Message,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("%v:%v", "localhost", "3008"))
	if err != nil {
		logrus.Errorf("Error listening %v", err)
	}

	logrus.Debugf("Server listening on port 3008")

	grpcServer := grpc.NewServer()
	pb.RegisterEchoServer(grpcServer, &EchoServer{})
	grpcServer.Serve(lis)
}
