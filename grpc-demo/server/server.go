package main

import (
	pb "GoProgrammingTourBook/grpc-demo/proto"
	"flag"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

var port string

func init() {
	flag.StringVar(&port, "p", "8000", "启动端口号")
	flag.Parse()
}

type GretterServer struct {
}

func (s *GretterServer) SayRoute(stream pb.Greeter_SayRouteServer) error {
	n := 0
	for {
		_ = stream.Send(&pb.HelloReply{Message: "say.route"})

		resp, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		n++
		log.Printf("resp: %v", resp)
	}
}

func main() {
	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, &GretterServer{})
	lis, _ := net.Listen("tcp", ":"+port)
	server.Serve(lis)
}
