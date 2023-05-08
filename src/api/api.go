package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	pb "github.com/tomoya_kamaji/go-pkg/grpc"
	"github.com/tomoya_kamaji/go-pkg/src/adapter/gorm"
	domain "github.com/tomoya_kamaji/go-pkg/src/domain/player"
	"github.com/tomoya_kamaji/go-pkg/src/infrastructure/repositoryImpl"
	usecase "github.com/tomoya_kamaji/go-pkg/src/usecase/player"
	"github.com/tomoya_kamaji/go-pkg/src/util/cast"
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

// 以下のスクレイピングコードは、別のパッケージに切り出すべき
var teamSiteUIDMap = map[string]string{
	"巨人":     "g",
	"阪神":     "t",
	"広島":     "c",
	"ヤクルト":   "s",
	"中日":     "d",
	"DeNA":   "yb",
	"西武":     "l",
	"ソフトバンク": "h",
	"楽天":     "e",
	"ロッテ":    "m",
	"オリックス":  "bs",
	"日本ハム":   "f",
}

func (s *baseBallApiServer) Crawler(ctx context.Context, in *emptypb.Empty) (*emptypb.Empty, error) {
	s.logger.Info(ctx, "start crawler")
	// ページ番号を指定して、URLを構築する

	teamPlayersMap := map[string][]domain.CreatePlayerParam{}
	for teamName, siteUID := range teamSiteUIDMap {
		fmt.Printf("\033[31m%#v\033[0m\n", teamName)
		url := fmt.Sprintf("https://baseball-data.com/stats/hitter-%v/tpa-1.html", siteUID)
		res, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		defer res.Body.Close()

		doc, err := goquery.NewDocumentFromReader(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		params := []domain.CreatePlayerParam{}
		doc.Find("#tbl tbody tr").Each(func(_ int, s *goquery.Selection) {
			param := domain.CreatePlayerParam{
				UniformNumber: cast.StringToInt64(s.Find("td").Eq(0).Text()),
				Name:          strings.ReplaceAll(s.Find("td").Eq(1).Text(), "\u3000", ""),
				AtBats:        cast.StringToInt64(s.Find("td").Eq(5).Text()),
				Hits:          cast.StringToInt64(s.Find("td").Eq(6).Text()),
				Walks:         cast.StringToInt64(s.Find("td").Eq(10).Text()),
				HomeRuns:      cast.StringToInt64(s.Find("td").Eq(7).Text()),
				RunsBattedIn:  cast.StringToInt64(s.Find("td").Eq(8).Text()),
			}
			params = append(params, param)
		})
		teamPlayersMap[teamName] = params
	}

	fmt.Printf("\033[31m%#v\033[0m\n", teamPlayersMap)
	return &emptypb.Empty{}, nil
}
