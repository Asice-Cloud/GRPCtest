package main

import (
	pb "GRPCtest/server/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	var remoteAddress string = "127.0.0.1:9876"
	connect, err := grpc.Dial(remoteAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect remote server %v , err is %v", remoteAddress, err)
	}
	defer connect.Close()

	//build connection
	client := pb.NewGraceClient(connect)
	//call
	var username string
	var age int
	fmt.Println("Please input username and age")
	fmt.Scanf("%s %d", &username, &age)
	resp, ok := client.Welcome(context.Background(), &pb.WelcomeRequest{
		RequestName: username,
		People: &pb.WelcomeRequestPerson{
			Name: username,
			Age:  int32(age),
		},
	})
	if ok != nil {
		log.Fatalf("Failed to call, err is ", ok)
	}
	fmt.Println(resp.GetResponseMessage())
}
