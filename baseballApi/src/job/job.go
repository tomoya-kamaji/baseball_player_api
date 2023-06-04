package job

import (
	"context"
	"fmt"

	"github.com/tomoya_kamaji/go-pkg/src/adapter/gorm"
	"github.com/tomoya_kamaji/go-pkg/src/infrastructure/repositoryImpl"
	usecase "github.com/tomoya_kamaji/go-pkg/src/usecase/player"
	"github.com/tomoya_kamaji/go-pkg/src/util/logging"
	"github.com/urfave/cli"
)

type jobServer struct {
	db     *gorm.DBProvider
	logger logging.Logger
}

func NewJobServer(db *gorm.DBProvider, logger logging.Logger) *jobServer {
	return &jobServer{db: db, logger: logger}
}

func (s *jobServer) Crawler(c *cli.Context) error {
	ctx := context.Background()
	player, err := usecase.NewFetchPlayerUsecase(repositoryImpl.NewPlayerRepositoryImpl(s.db)).Run(ctx, "1")
	if err != nil {
		return err
	}
	// 出力
	fmt.Printf("\033[31m%#v\033[0m\n", player)
	return nil
}
