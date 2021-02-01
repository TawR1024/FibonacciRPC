package main

import (
	"context"
	"fmt"
	pb "github.com/TawR1024/FibonacciRPC/fibonaccirpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func main(){
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	conn, err := grpc.Dial("127.0.0.1:3000", opts...)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewFibonacciClient(conn)

	request := &pb.Request{
		Position: 1,
	}
	response, err := client.CalculateFibonacci(context.Background(), request)
	if err !=nil{
		grpclog.Fatalf("Fail to dial: %v", err)
	}
	fmt.Println("Fibonnacci number is:", response.Fnumber)


	request = &pb.Request{
		Position: 10,
	}
	response, err = client.CalculateFibonacci(context.Background(), request)
	if err !=nil{
		grpclog.Fatalf("Fail to dial: %v", err)
	}
	fmt.Println("Fibonnacci number is:", response.Fnumber)



	request = &pb.Request{
		Position: 33,
	}
	response, err = client.CalculateFibonacci(context.Background(), request)
	if err !=nil{
		grpclog.Fatalf("Fail to dial: %v", err)
	}
	fmt.Println("Fibonnacci number is:", response.Fnumber)
}