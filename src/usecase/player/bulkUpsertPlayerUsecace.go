package usecase

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	domain "github.com/tomoya_kamaji/go-pkg/src/domain/player"
	"github.com/tomoya_kamaji/go-pkg/src/usecase/transaction"
	"github.com/tomoya_kamaji/go-pkg/src/util/cast"
)

type BulkUpsertPlayerPlayerUsecase struct {
	txMgr            transaction.TransactionManager
	playerRepository domain.PlayerRepository
}

func NewBulkUpsertPlayerPlayerUsecase(txMgr transaction.TransactionManager, playerRepository domain.PlayerRepository) *BulkUpsertPlayerPlayerUsecase {
	return &BulkUpsertPlayerPlayerUsecase{
		txMgr:            txMgr,
		playerRepository: playerRepository,
	}
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

func (u *BulkUpsertPlayerPlayerUsecase) Run(ctx context.Context) error {
	start := time.Now()
	defer func() {
		fmt.Printf("=======経過時間: %f秒=======\n", time.Since(start).Seconds())
	}()

	for teamName, siteUID := range teamSiteUIDMap {
		fmt.Println(teamName)
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

		players := []*domain.Player{}
		doc.Find("#tbl tbody tr").Each(func(_ int, s *goquery.Selection) {
			player := domain.CreatePlayer(
				domain.CreatePlayerParam{
					UniformNumber: cast.StringToInt64(s.Find("td").Eq(0).Text()),
					Name:          strings.ReplaceAll(s.Find("td").Eq(1).Text(), "\u3000", ""),
					AtBats:        cast.StringToInt64(s.Find("td").Eq(5).Text()),
					Hits:          cast.StringToInt64(s.Find("td").Eq(6).Text()),
					Walks:         cast.StringToInt64(s.Find("td").Eq(10).Text()),
					HomeRuns:      cast.StringToInt64(s.Find("td").Eq(7).Text()),
					RunsBattedIn:  cast.StringToInt64(s.Find("td").Eq(8).Text()),
				},
			)
			players = append(players, player)
		})
		u.playerRepository.BulkUpsert(ctx, players)
	}
	return nil
}

