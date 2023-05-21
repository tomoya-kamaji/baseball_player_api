package queryImpl

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	domain "github.com/tomoya_kamaji/go-pkg/src/domain/player"
	"github.com/tomoya_kamaji/go-pkg/src/infrastructure/repositoryImpl"
	"github.com/tomoya_kamaji/go-pkg/src/query"
	"github.com/tomoya_kamaji/go-pkg/src/util/testutil"
)

func TestPlayerSearchQueryImpl(t *testing.T) {
	// desc := query.DESC
	asc := query.ASC
	hits := query.Hits
	// テストケースの定義
	tests := []struct {
		name           string
		input          []*domain.Player
		searchParam    query.SearchParam
		expectedResult []string
	}{
		{
			name: "検索条件なし",
			input: []*domain.Player{
				hits100,
				hits99,
			},
			searchParam: query.SearchParam{},
			expectedResult: []string{
				hits100.Name,
				hits99.Name,
			},
		},
		{
			name: "Hits 100以上の検索",
			input: []*domain.Player{
				hits100,
				hits99,
			},
			searchParam: query.SearchParam{
				MinHits: 100,
			},
			expectedResult: []string{
				hits100.Name,
			},
		},
		{
			name: "Hits 99未満の検索",
			input: []*domain.Player{
				hits100,
				hits99,
			},
			searchParam: query.SearchParam{
				MaxHits: 99,
			},
			expectedResult: []string{
				hits99.Name,
			},
		},
		{
			name: "Hitsでsort",
			input: []*domain.Player{
				hits100,
				hits99,
				hits98,
			},
			searchParam: query.SearchParam{
				SortField: &hits,
				SortOrder: &asc,
			},
			expectedResult: []string{
				hits100.Name,
				hits99.Name,
				hits98.Name,
			},
		},
	}

	db := testutil.NewTestMainDB()
	queryImpl := NewPlayerSearchQueryImpl(db)
	repositoryImpl := repositoryImpl.NewPlayerRepositoryImpl(db)

	for _, td := range tests {
		t.Run(td.name, func(t *testing.T) {
			ctx := context.Background()
			repositoryImpl.BulkUpsert(ctx, td.input)

			// 検索
			dto, err := queryImpl.Run(ctx, td.searchParam)
			assert.NoError(t, err)

			names := []string{}
			for _, player := range dto.Players {
				names = append(names, player.Name)
			}
			// assert
			assert.ElementsMatch(t, names, td.expectedResult)
		})
	}

}

var hits100 = domain.CreatePlayer(
	domain.CreatePlayerParam{
		Name: "100本ヒッター",
		Hits: 100,
	},
)
var hits99 = domain.CreatePlayer(
	domain.CreatePlayerParam{
		Name: "99本ヒッター",
		Hits: 99,
	},
)
var hits98 = domain.CreatePlayer(
	domain.CreatePlayerParam{
		Name: "98本ヒッター",
		Hits: 98,
	},
)
