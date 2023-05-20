package api

import (
	"context"

	pb "github.com/tomoya_kamaji/go-pkg/grpc"
	"github.com/tomoya_kamaji/go-pkg/src/infrastructure/repositoryImpl"
	usecase "github.com/tomoya_kamaji/go-pkg/src/usecase/player"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *baseBallApiServer) CreatePlayer(ctx context.Context, in *pb.CreatePlayersRequest) (*pb.CreatePlayerResponse, error) {
	param := usecase.CreatePlayerUsecaseParam{
		UniformNumber: in.GetUniformNumber(),
		Name:          in.GetName(),
		AtBats:        in.GetAtBats(),
		Hits:          in.GetHits(),
		Walks:         in.GetWalks(),
		HomeRuns:      in.GetHomeRuns(),
		RunsBattedIn:  in.GetRunsBattedIn(),
	}
	dto, err := usecase.NewCreatePlayerUsecase(
		repositoryImpl.NewTransactionManagerImpl(s.db),
		repositoryImpl.NewPlayerRepositoryImpl(s.db),
	).Run(ctx, param)

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.CreatePlayerResponse{Player: &pb.Player{
		Id:            dto.ID,
		UniformNumber: dto.UniformNumber,
		Name:          dto.Name,
		AtBats:        dto.AtBats,
		Hits:          dto.Hits,
		Walks:         dto.Walks,
		HomeRuns:      dto.HomeRuns,
		RunsBattedIn:  dto.RunsBattedIn,
	}}, nil
}
