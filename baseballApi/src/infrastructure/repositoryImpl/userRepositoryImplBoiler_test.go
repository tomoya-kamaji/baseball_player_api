package repositoryImpl

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	domain "github.com/tomoya_kamaji/go-pkg/src/domain/user"
	"github.com/tomoya_kamaji/go-pkg/src/util/testutil"
)

func TestUserRepository_GetById(t *testing.T) {
	db, _ := testutil.NewSqlBoilerMainDB()
	tests := []struct {
		name string
		id   domain.UserID
	}{
		{
			name: "プレイヤーが作成されること",
			id:   domain.UserID("1"),
		},
	}
	for _, td := range tests {
		t.Run(fmt.Sprintf(": %s", td.name), func(t *testing.T) {
			ctx := context.Background()

			repository := NewUserRepositoryBoilerImpl(db)
			u, err := repository.GetByID(ctx, td.id)
			if err != nil {
				t.Fatalf("error: %s", err)
			}
			assert.Equal(t, td.id, u.ID)
		})
	}
}
