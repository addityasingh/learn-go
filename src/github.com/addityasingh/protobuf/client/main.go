package main

import (
	"context"
	"fmt"

	"github.com/Sirupsen/logrus"
	pb "github.com/addityasingh/protobuf/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:3008", grpc.WithInsecure())

	if err != nil {
		logrus.Errorf("failure while dialing: %v", err)
	}
	defer conn.Close()

	client := pb.NewEchoClient(conn)
	response, err := client.Reply(context.Background(), &pb.EchoRequest{
		Id:      "1",
		Message: "Hello World!",
	})
	fmt.Println(response)
}
