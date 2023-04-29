package repositoryImpl

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	domain "github.com/tomoya_kamaji/go-pkg/src/domain/player"
	"github.com/tomoya_kamaji/go-pkg/src/util/testutil"
)

func TestUserRepository_Create(t *testing.T) {
	ctx := context.Background()
	repository := NewPlayerRepositoryImpl(testutil.NewTestMainDB())

	player := domain.CreatePlayer(
		domain.CreatePlayerParam{
			UniformNumber: 51,
			Name:          "イチロー",
			AtBats:        550,
			Hits:          200,
			HomeRuns:      10,
			RunsBattedIn:  70,
			Walks:         50,
		},
	)
	repository.Create(ctx, player)
	result, _ := repository.GetByID(ctx, player.ID)
	expected := player

	assert.Equal(t, result, expected)
}
