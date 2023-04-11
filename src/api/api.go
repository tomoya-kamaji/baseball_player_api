package api

import (
	"context"

	pb "github.com/tomoya_kamaji/go-pkg/grpc"
)

type baseBallApiServer struct {
	pb.UnimplementedBaseBallApiServer
}

func NewApiServer() pb.BaseBallApiServer {
	return &baseBallApiServer{}
}

func (s *baseBallApiServer) SelectPlayers(ctx context.Context, in *pb.SelectPlayersRequest) (*pb.SelectPlayersResponse, error) {
	player1 := &pb.Player{
		Id:       1,
		Name:     "Tomoya",
		Position: "Pitcher",
	}

	player2 := &pb.Player{
		Id:       2,
		Name:     "Tomoya",
		Position: "Pitcher",
	}
	players := []*pb.Player{player1, player2}
	return &pb.SelectPlayersResponse{Players: players}, nil
}
