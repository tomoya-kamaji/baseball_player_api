package api

import (
	"context"
	"fmt"
	"time"

	"github.com/tomoya_kamaji/go-pkg/src/infrastructure/repositoryImpl"
	usecase "github.com/tomoya_kamaji/go-pkg/src/usecase/player"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *baseBallApiServer) Crawler(ctx context.Context, in *emptypb.Empty) (*emptypb.Empty, error) {
	s.logger.Info(ctx, "start crawler")
	start := time.Now()
	defer func() {
		fmt.Printf("=======経過時間: %f秒=======\n", time.Since(start).Seconds())
	}()

	err := usecase.NewCrawlPlayerUsecase(
		repositoryImpl.NewTransactionManagerImpl(s.db),
		repositoryImpl.NewPlayerRepositoryImpl(s.db),
	).Run(ctx)

	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &emptypb.Empty{}, nil
}
