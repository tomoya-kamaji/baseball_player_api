package repositoryImpl

import (
	"context"
	"database/sql"
	"fmt"

	domain "github.com/tomoya_kamaji/go-pkg/src/domain/user"
	"github.com/tomoya_kamaji/go-pkg/src/infrastructure/sqlboiler/models"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type userRepositoryBoilerImpl struct {
	db *sql.DB
}

func NewUserRepositoryBoilerImpl(db *sql.DB) domain.UserRepository {
	return &userRepositoryBoilerImpl{db: db}
}

func (r *userRepositoryBoilerImpl) Save(c context.Context, u *domain.User) error {
	user := &models.User{
		ID:            u.ID.String(),
		Name:          u.Name,
		Address:       u.Address,
		ContactNumber: u.ContactNumber,
	}
	err := user.Insert(c, r.db, boil.Columns{})
	if err != nil {
		return err
	}
	return nil
}

func (r *userRepositoryBoilerImpl) GetByID(c context.Context, id domain.UserID) (*domain.User, error) {
	// primary keyでfind
	user, err := models.FindUser(c, r.db, id.String())
	if err != nil {
		return nil, err
	}

	// 動的なクエリ
	var name string
	queryMods := []qm.QueryMod{
		models.UserWhere.ID.EQ(id.String()),
	}
	// nameが空でない場合は、クエリに追加
	if name != "" {
		queryMods = append(queryMods, qm.Where("name LIKE ?", "%"+name+"%"))
	}
	users, err := models.Users(queryMods...).All(c, r.db)
	fmt.Printf("users: %v\n", users)
	return r.convertToModel(user), nil
}

func (repo *userRepositoryBoilerImpl) convertToModel(u *models.User) *domain.User {
	return domain.ReConstractUser(
		domain.ReConstractUserParam{
			ID:            domain.UserID(u.ID),
			Name:          u.Name,
			Address:       u.Address,
			ContactNumber: u.ContactNumber,
		},
	)
}
