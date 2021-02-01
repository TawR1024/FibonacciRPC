package main

import (
	"context"
	"fmt"
	pb "github.com/TawR1024/FibonacciRPC/fibonaccirpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"math"
	"net"
)
type FibonacciServer struct {
	pb.UnimplementedFibonacciServer
}

var sqrtOfFive = math.Sqrt(5)

func (s *FibonacciServer)CalculateFibonacci(ctx context.Context, request *pb.Request)(response *pb.Reply, err error)  {
	f := (math.Pow((1+sqrtOfFive)/2, float64(request.Position)) - math.Pow((1-sqrtOfFive)/2, float64(request.Position))) / sqrtOfFive
	response = &pb.Reply{Fnumber: fmt.Sprintf("%f", f)}
	return response, nil

}

func main(){
	listener, err := net.Listen("tcp", ":3000")
	if err !=nil{
		grpclog.Fatalf("Failed to listen: %v", err)
	}
	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterFibonacciServer(grpcServer, &FibonacciServer{})

	grpcServer.Serve(listener)

}
