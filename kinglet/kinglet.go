package kinglet

import (
	"context"
	"net"

	pb "github.com/import-yuefeng/kingBird/pb/test"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func init() {
}

const (
	PORT = ":2048"
)

type server struct{}

func Kinglet(configPath string) {
	log.Info("hello, now is kinglet mode")

	lis, err := net.Listen("tcp", PORT)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Info("start gRPC service...")
	s.Serve(lis)

}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Infoln("request: ", in.Name)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}
