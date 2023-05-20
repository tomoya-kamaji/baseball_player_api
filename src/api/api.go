package api

import (
	pb "github.com/tomoya_kamaji/go-pkg/grpc"
	"github.com/tomoya_kamaji/go-pkg/src/adapter/gorm"
	"github.com/tomoya_kamaji/go-pkg/src/util/logging"
)

type baseBallApiServer struct {
	db     *gorm.DBProvider
	logger logging.Logger
	pb.UnimplementedBaseBallApiServer
}

func NewApiServer(db *gorm.DBProvider, logger logging.Logger) pb.BaseBallApiServer {
	return &baseBallApiServer{db: db, logger: logger}
}
