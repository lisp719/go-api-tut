package user

import (
	"context"
	"log"
	"time"

	pb "go-api-tut/pkg/proto"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func Hello(c *gin.Context) {
	conn, err := grpc.Dial("grpc:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := client.SayHello(ctx, &pb.HelloRequest{Name: "world"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	c.JSON(200, gin.H{"message": r.GetMessage()})
}
