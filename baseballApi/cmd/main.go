package main

import (
	"github.com/tomoya_kamaji/go-pkg/src/config"
	"github.com/tomoya_kamaji/go-pkg/src/route"
	"github.com/tomoya_kamaji/go-pkg/src/infrastructure/sendgrid"
	v1 "github.com/tomoya_kamaji/go-pkg/src/route/v1"
)

const (
	port = ":8082"
)

// @title Baseball API
// @version バージョン(1.0)
// @description 野球選手の成績を管理するAPIを提供する
// @license.name ライセンス(必須)
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8082
func main() {
	api := route.NewEngine()
	config.InitLogger()

	sendgrid.InitNotificationWorker()
	v1.Init(api)
	api.Run(port)
}

// gRPCサーバーを起動する
// func main() {
// // listenPort
// listenPort, err := net.Listen("tcp", port)
// if err != nil {
// 	panic(err)
// }

// // gRPCサーバーを作成
// server := grpc.NewServer()
// baseBallApi := api.NewApiServer(gorm.NewMainDB(), logging.NewStackDriverLoggerByName("baseball-api"))
// pb.RegisterBaseBallApiServer(server, baseBallApi)
// reflection.Register(server)
// go func() {
// 	log.Printf("start gRPC server port: %v", port)
// 	err = server.Serve(listenPort)
// 	if err != nil && err != grpc.ErrServerStopped {
// 		panic(err)
// 	}
// }()

// // Ctrl+Cが入力されたらGraceful shutdownされるようにする
// sig := make(chan os.Signal, 1)
// signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
// fmt.Printf("signal received: %s\n", <-sig)
// server.GracefulStop()
// }
