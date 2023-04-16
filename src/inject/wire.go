package inject

import (
	"github.com/google/wire"
	pb "github.com/tomoya_kamaji/go-pkg/grpc"
	"github.com/tomoya_kamaji/go-pkg/src/adapter/gorm"
	"github.com/tomoya_kamaji/go-pkg/src/api"
)

func InitializeBaseBallApiServer() pb.BaseBallApiServer {
	dbProvider := gorm.NewMainDB()
	wire.Build(wire.NewSet(
		api.NewApiServer,
		dbProvider,
	))
	return nil
}
