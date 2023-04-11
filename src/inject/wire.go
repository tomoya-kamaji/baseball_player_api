//go:build wireinject

package inject

import (
	"github.com/google/wire"
	pb "github.com/tomoya_kamaji/go-pkg/grpc"
	"github.com/tomoya_kamaji/go-pkg/src/api"
)

func InitializeBaseBallApiServer() pb.BaseBallApiServer {
	wire.Build(wire.NewSet(
		api.NewApiServer,
	))
	return nil
}
