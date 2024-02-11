package main

import (
	pb "GRPCtest/server/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

func main() {
	//ssl authorization:
	creds, credErr := credentials.NewServerTLSFromFile("key/test.pem", "*")
	if credErr != nil {
		log.Fatalf("License failure,error is: %v", credErr)
	}

	//verifier
	var remoteAddress string = "127.0.0.1:9999"
	connect, err := grpc.Dial(remoteAddress, grpc.WithTransportCredentials(creds))
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
