package kinglet

import (
	"context"
	"net"

	task "github.com/import-yuefeng/kingBird/pb/task"
	test "github.com/import-yuefeng/kingBird/pb/test"
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
	test.RegisterGreeterServer(s, &server{})
	task.RegisterTaskerServer(s, &server{})
	log.Info("start gRPC service...")
	s.Serve(lis)

}

func (s *server) SayHello(ctx context.Context, in *test.HelloRequest) (*test.HelloReply, error) {
	log.Infoln("request: ", in.Name)
	return &test.HelloReply{Message: "Hello " + in.Name}, nil
}

func (s *server) RunNewTask(ctx context.Context, in *task.RunTaskRequest) (*task.TaskReply, error) {
	log.Infoln("request: ", in)
	return &task.TaskReply{IsSuccess: true}, nil
}

func (s *server) StopTask(ctx context.Context, in *task.StopTaskRequest) (*task.TaskReply, error) {
	log.Infoln("request: ", in)
	return &task.TaskReply{IsSuccess: true}, nil
}

func (s *server) ConitnueTask(ctx context.Context, in *task.ContinueTaskRequest) (*task.TaskReply, error) {
	log.Infoln("request: ", in)
	return &task.TaskReply{IsSuccess: true}, nil
}
func (s *server) CheckTask(ctx context.Context, in *task.CheckTaskRequest) (*task.TaskReply, error) {
	log.Infoln("request: ", in)
	return &task.TaskReply{IsSuccess: true}, nil
}
func (s *server) SaveTask(ctx context.Context, in *task.SaveTaskRequest) (*task.TaskReply, error) {
	log.Infoln("request: ", in)
	return &task.TaskReply{IsSuccess: true}, nil
}
