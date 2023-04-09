package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	pb "github.com/tomoya_kamaji/go-pkg/grpc"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func NewApiServer() *apiServer {
	return &apiServer{}
}

type apiServer struct {
	pb.UnimplementedBaseBallApiServer
}

func (s *apiServer) SelectPlayers(ctx context.Context, in *pb.SelectPlayersRequest) (*pb.SelectPlayersResponse, error) {
	player1 := &pb.Player{
		Id:       1,
		Name:     "Tomoya",
		Position: "Pitcher",
	}

	player2 := &pb.Player{
		Id:       1,
		Name:     "Tomoya",
		Position: "Pitcher",
	}

	players := []*pb.Player{player1, player2}

	return &pb.SelectPlayersResponse{Players: players}, nil
}
func main() {
	// listenPortを作成
	listenPort, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	// gRPCサーバーを作成
	server := grpc.NewServer()
	pb.RegisterBaseBallApiServer(server, NewApiServer())
	go func() {
		log.Printf("start gRPC server port: %v", port)
		err = server.Serve(listenPort)
		if err != nil && err != grpc.ErrServerStopped {
			panic(err)
		}
	}()

	// Ctrl+Cが入力されたらGraceful shutdownされるようにする
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("signal received: %s\n", <-sig)
	server.GracefulStop()
}
