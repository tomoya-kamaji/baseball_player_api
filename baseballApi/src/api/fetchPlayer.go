package api

import (
	"context"

	pb "github.com/tomoya_kamaji/go-pkg/grpc"
	"github.com/tomoya_kamaji/go-pkg/src/infrastructure/repositoryImpl"
	usecase "github.com/tomoya_kamaji/go-pkg/src/usecase/player"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *baseBallApiServer) FetchPlayer(ctx context.Context, in *pb.FetchPlayerRequest) (*pb.FetchPlayerResponse, error) {
	playerId := in.GetPlayerId()
	player, err := usecase.NewFetchPlayerUsecase(repositoryImpl.NewPlayerRepositoryImpl(s.db)).Run(ctx, playerId)

	if err != nil {
		s.logger.Error(ctx, "failed to fetch player", zap.Error(err))
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	return &pb.FetchPlayerResponse{Player: &pb.Player{Id: string(player.ID), Name: player.Name}}, nil
}
