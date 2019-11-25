package kingctl

import (
	config "github.com/import-yuefeng/kingBird/config"
	pb "github.com/import-yuefeng/kingBird/pb/test"
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
	config.ReadConfigFile(configPath)

	c := pb.NewGreeterClient(conn)

	name := "kingBird"

	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Println(r.Message)
}
