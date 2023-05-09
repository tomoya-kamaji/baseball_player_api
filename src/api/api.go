package api

import (
	"context"
	"fmt"
	"time"

	pb "github.com/tomoya_kamaji/go-pkg/grpc"
	"github.com/tomoya_kamaji/go-pkg/src/adapter/gorm"
	"github.com/tomoya_kamaji/go-pkg/src/infrastructure/repositoryImpl"
	usecase "github.com/tomoya_kamaji/go-pkg/src/usecase/player"
	"github.com/tomoya_kamaji/go-pkg/src/util/logging"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type baseBallApiServer struct {
	db     *gorm.DBProvider
	logger logging.Logger
	pb.UnimplementedBaseBallApiServer
}

func NewApiServer(db *gorm.DBProvider, logger logging.Logger) pb.BaseBallApiServer {
	return &baseBallApiServer{db: db, logger: logger}
}

func (s *baseBallApiServer) FetchPlayer(ctx context.Context, in *pb.FetchPlayerRequest) (*pb.FetchPlayerResponse, error) {
	playerId := in.GetPlayerId()
	player, err := usecase.NewFetchPlayerUsecase(repositoryImpl.NewPlayerRepositoryImpl(s.db)).Run(ctx, playerId)

	if err != nil {
		s.logger.Error(ctx, "failed to fetch player", zap.Error(err))
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	return &pb.FetchPlayerResponse{Player: &pb.Player{Id: string(player.ID), Name: player.Name}}, nil
}

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

func (s *baseBallApiServer) Crawler(ctx context.Context, in *emptypb.Empty) (*emptypb.Empty, error) {
	s.logger.Info(ctx, "start crawler")
	start := time.Now()
	defer func() {
		fmt.Printf("=======経過時間: %f秒=======\n", time.Since(start).Seconds())
	}()

	err := usecase.NewBulkUpsertPlayerPlayerUsecase(
		repositoryImpl.NewTransactionManagerImpl(s.db),
		repositoryImpl.NewPlayerRepositoryImpl(s.db),
	).Run(ctx)

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}
