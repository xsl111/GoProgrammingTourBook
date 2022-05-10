package main

import (
	pb "GoProgrammingTourBook/grpc-demo/proto"
	"context"
	"flag"
	"io"
	"log"

	"google.golang.org/grpc"
)

var port string

func init() {
	flag.StringVar(&port, "p", "8000", "启动端口号")
	flag.Parse()
}

func main() {
	conn, _ := grpc.Dial(":"+port, grpc.WithInsecure())
	defer conn.Close()

	client := pb.NewGreeterClient(conn)
	_ = SayRoute(client, &pb.HelloRequest{Name: "许仕蕾你好啊"})
}

func SayRoute(client pb.GreeterClient, r *pb.HelloRequest) error {
	stream, _ := client.SayRoute(context.Background())
	for n := 0; n <= 6; n++ {
		_ = stream.Send(r)
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		log.Printf("resp err: %v", resp)
	}
	_ = stream.CloseSend()
	return nil
}
