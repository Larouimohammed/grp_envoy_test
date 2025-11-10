package main

import (
    "fmt"
    "log"
    "net"
    "time"

   pb "github.com/yourusername/grpc_envoy_test/helloworld"

    "google.golang.org/grpc"
)

type server struct {
    pb.UnimplementedGreeterServer
}

func (s *server) SayHello(req *pb.HelloRequest, stream pb.Greeter_LotsOfRepliesServer) error {
    for i := 1; i <= 5; i++ {
        msg := fmt.Sprintf("Hello %s, message %d", req.Name, i)
        if err := stream.Send(&pb.HelloReply{Message: msg}); err != nil {
            return err
        }
        time.Sleep(time.Second)
    }
    return nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    pb.RegisterGreeterServer(s, &server{})
    log.Println("gRPC streaming server listening on :50051")
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}

