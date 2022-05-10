package main

import (
	pb "GoProgrammingTourBook/tag-service/proto"
	"GoProgrammingTourBook/tag-service/server"
	"flag"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	var port string
	s := grpc.NewServer()
	pb.RegisterTagServiceServer(s, server.NewTagServer())
	reflection.Register(s)
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("net.listen err: %v", err)
	}

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("server.serve err: %v", err)
	}
	flag.StringVar(&port, "port", "8000", "端口号")
}
