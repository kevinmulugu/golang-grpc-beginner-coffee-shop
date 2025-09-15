package main

import (
	"context"
	"log"
	"net"

	pb "github.com/kevinmulugu/proto_example/coffeshop_protos"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedCoffeeShopServer
}

func (s *Server) GetMenu(menuRequest *pb.MenuRequest, srv grpc.ServerStreamingServer[pb.Menu]) error {
	items := []*pb.Item{
		{
			Id:   "1",
			Name: "Mandazi",
		},
		{
			Id:   "2",
			Name: "Chapati",
		},
		{
			Id:   "3",
			Name: "Mayai Chafua",
		},
		{
			Id:   "4",
			Name: "Ndizi",
		},
	}

	for i := range items {
		srv.Send(&pb.Menu{
			Items: items[0 : i+1],
		})
	}

	return nil
}

func (s *Server) PlaceOrder(ctx context.Context, order *pb.Order) (*pb.Receipt, error) {
	return &pb.Receipt{
		Id: "Qwerty",
	}, nil
}

func (s *Server) GetOrderStatus(ctx context.Context, receipt *pb.Receipt) (*pb.OrderStatus, error) {
	return &pb.OrderStatus{
		OrderId: receipt.Id,
		Status:  "IN PROGRESS",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9001");
	if err != nil {
		log.Fatalf("failed to listen %s\n", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCoffeeShopServer(grpcServer, &Server{})

	log.Println("Coffee shop gRPC server is listening on :9001")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve :%s\n", err)
	}
}
