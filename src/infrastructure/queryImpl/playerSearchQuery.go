package queryimpl

import (
	"context"
	"fmt"

	"github.com/tomoya_kamaji/go-pkg/src/adapter/gorm"
	"github.com/tomoya_kamaji/go-pkg/src/infrastructure/entity"
	"github.com/tomoya_kamaji/go-pkg/src/query"
)

type playerSearchQueryImpl struct {
	db *gorm.DBProvider
}

func NewPlayerSearchQueryImpl(db *gorm.DBProvider) query.PlayerSearchQuery {
	return &playerSearchQueryImpl{db: db}
}

func (q *playerSearchQueryImpl) Run(c context.Context, param query.SearchParam) (*query.PlayerSearchQueryDTO, error) {
	dbQuery := q.db.Provide(c).Model(&entity.PlayerEntity{})
	// フィルター条件の設定
	if param.MinHits > 0 {
		dbQuery = dbQuery.Where("hits >= ?", param.MinHits)
	}
	if param.MaxHits > 0 {
		dbQuery = dbQuery.Where("hits <= ?", param.MaxHits)
	}
	if param.MinHomeRuns > 0 {
		dbQuery = dbQuery.Where("home_runs >= ?", param.MinHomeRuns)
	}
	if param.MaxHomeRuns > 0 {
		dbQuery = dbQuery.Where("home_runs <= ?", param.MaxHomeRuns)
	}
	if param.MinRunsBattedIn > 0 {
		dbQuery = dbQuery.Where("runs_batted_in >= ?", param.MinRunsBattedIn)
	}
	if param.MaxRunsBattedIn > 0 {
		dbQuery = dbQuery.Where("runs_batted_in <= ?", param.MaxRunsBattedIn)
	}

	// ソート設定
	if param.SortField != "" && param.SortOrder != "" {
		sortOrder := "asc"
		if param.SortOrder == "desc" {
			sortOrder = "desc"
		}
		dbQuery = dbQuery.Order(fmt.Sprintf("%s %s", param.SortField, sortOrder))
	}

	// 取得件数制限
	if param.Limit > 0 {
		dbQuery = dbQuery.Limit(param.Limit)
	}

	var players []*query.PlayerDTO
	if err := dbQuery.Find(&players).Error; err != nil {
		return nil, err
	}

	return &query.PlayerSearchQueryDTO{
		Players: players,
	}, nil
}
