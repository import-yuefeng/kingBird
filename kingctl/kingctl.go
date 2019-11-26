package kingctl

import (
	config "github.com/import-yuefeng/kingBird/config"
	task "github.com/import-yuefeng/kingBird/pb/task"
	test "github.com/import-yuefeng/kingBird/pb/test"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = ":2048"
)

func init() {

}

func Kingctl(configPath string) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	_, err = config.ReadConfigFile(configPath)
	// log.Infoln(res)
	if err != nil {
		log.Fatalln(err)
		return
	}
	sayHello(conn)
	runTask(conn, "ls -hal", "host", "config....")
}

func sayHello(conn *grpc.ClientConn) {
	c := test.NewGreeterClient(conn)
	name := "kingBird"
	r, err := c.SayHello(context.Background(), &test.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Println(r.Message)
}

func runTask(conn *grpc.ClientConn, cmd string, mode string, config string) {
	c := task.NewTaskerClient(conn)
	taskInfo := &task.RunTaskRequest{
		Cmd:    cmd,
		Mode:   mode,
		Config: config,
	}

	r, err := c.RunNewTask(context.Background(), taskInfo)
	if err != nil {
		log.Fatalf("could not send request with: %v", err)
	}
	log.Println(r.IsSuccess)
}
