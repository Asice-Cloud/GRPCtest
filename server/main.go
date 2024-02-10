package main

import (
	pb "GRPCtest/server/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedGraceServer
}

func (s *server) Welcome(ctx context.Context, req *pb.WelcomeRequest) (*pb.WelcomeResponse, error) {
	fmt.Println("Find user: " + req.RequestName)
	return &pb.WelcomeResponse{
		ResponseMessage: "hello, user:" + req.RequestName + "\nhappy birthday for your " + fmt.Sprintf("%d", req.People.Age)}, nil
}

func main() {
	//Authorization:

	//1.open port:
	listen, err := net.Listen("tcp", ":9876")
	if err != nil {
		log.Fatalf("Failed to open port,err is %v", err)
		return
	}
	//2.create grpc service:
	grpcServer := grpc.NewServer()
	//3. register our service:
	pb.RegisterGraceServer(grpcServer, &server{})
	//4.launch server:
	launchErr := grpcServer.Serve(listen)
	if launchErr != nil {
		log.Fatalf("Failed to launch grpc server, err is %v", err)
		return
	}
}
