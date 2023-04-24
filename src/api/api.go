package api

import (
	"context"

	pb "github.com/tomoya_kamaji/go-pkg/grpc"
	"github.com/tomoya_kamaji/go-pkg/src/adapter/gorm"
	"github.com/tomoya_kamaji/go-pkg/src/infrastructure/repositoryImpl"
	usecase "github.com/tomoya_kamaji/go-pkg/src/usecase/player"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type baseBallApiServer struct {
	db *gorm.DBProvider
	pb.UnimplementedBaseBallApiServer
}

func NewApiServer(db *gorm.DBProvider) pb.BaseBallApiServer {
	return &baseBallApiServer{db: db}
}

func (s *baseBallApiServer) SelectPlayers(ctx context.Context, in *pb.SelectPlayersRequest) (*pb.SelectPlayersResponse, error) {
	playerIds := in.GetPlayerIds()
	player, err := usecase.NewFetchPlayerUsecase(
		repositoryImpl.NewPlayerRepositoryImpl(s.db),
	).Run(ctx, playerIds[0])

	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	player1 := &pb.Player{
		Id:   string(player.ID),
		Name: player.Name,
	}

	players := []*pb.Player{player1}
	return &pb.SelectPlayersResponse{Players: players}, nil
}
