package main

import (
	pb "GRPCtest/server/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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
	creds, credErr := credentials.NewServerTLSFromFile("key/test.pem", "key/test.key")
	if credErr != nil {
		log.Fatalf("License error %v", credErr)
		return
	}

	//1.open port:
	listen, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatalf("Failed to open port,err is %v", err)
		return
	}
	//2.create grpc service:
	grpcServer := grpc.NewServer(grpc.Creds(creds))
	//3. register our service:
	pb.RegisterGraceServer(grpcServer, &server{})
	//4.launch server:
	launchErr := grpcServer.Serve(listen)
	if launchErr != nil {
		log.Fatalf("Failed to launch grpc server, err is %v", err)
		return
	}
}
