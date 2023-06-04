package repositoryImpl

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	domain "github.com/tomoya_kamaji/go-pkg/src/domain/player"
	"github.com/tomoya_kamaji/go-pkg/src/util/testutil"
)

func TestUserRepository_Create(t *testing.T) {
	tests := []struct {
		name  string
		input *domain.Player
	}{
		{
			name: "プレイヤーが作成されること",
			input: domain.CreatePlayer(
				domain.CreatePlayerParam{
					UniformNumber: 51,
					Name:          "イチロー",
					AtBats:        550,
					Hits:          200,
					HomeRuns:      10,
					RunsBattedIn:  70,
					Walks:         50,
				},
			),
		},
	}
	for _, td := range tests {
		t.Run(fmt.Sprintf(": %s", td.name), func(t *testing.T) {
			ctx := context.Background()
			repository := NewPlayerRepositoryImpl(testutil.NewTestMainDB())
			repository.Create(ctx, td.input)
			result, _ := repository.GetByID(ctx, td.input.ID)

			// assert
			assert.Equal(t, result, td.input)
		})
	}
}
