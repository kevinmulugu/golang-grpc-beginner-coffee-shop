package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/kevinmulugu/proto_example/coffeshop_protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("127.0.0.1:9001", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("failed to connect to gRPC server")
	}

	defer conn.Close()

	c := pb.NewCoffeeShopClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	menuStream, err := c.GetMenu(ctx, &pb.MenuRequest{})
	if err != nil {
		log.Fatalf("error  calling GetMenu: %v", err)
	}

	done := make(chan bool)
	var items []*pb.Item
	go func() {
		for {
			resp, err := menuStream.Recv()
			if err == io.EOF {
				done <- true
				return
			}

			if err != nil {
				log.Fatalf("can not receive %v", err)
			}

			items = resp.Items
			log.Printf("Resp received: %v", resp.Items)
		}
	}()

	<-done

	receipt, err := c.PlaceOrder(ctx, &pb.Order{Items: items})
	if err != nil {
		log.Fatalf("error placing order: %s", err.Error())
	}
	log.Printf("%v", receipt)

	status, err := c.GetOrderStatus(ctx, &pb.Receipt{Id: receipt.Id})
	if err != nil {
		log.Fatalf("error getting order status: %v", err)
	}
	log.Printf("order status: %v", status)
}
