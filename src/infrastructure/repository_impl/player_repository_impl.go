package repository

import (
	"context"

	"github.com/tomoya_kamaji/go-pkg/src/infrastructure/entity"
	"gorm.io/gorm"
)

type PlayerRepository struct {
	DB *gorm.DB
}

func (pr *PlayerRepository) GetByID(c context.Context, id string) (*entity.PlayerEntity, error) {
	var player entity.PlayerEntity
	// playerテーブルからidをキーにレコードを取得する
	if err := pr.DB.First(&player, id).Error; err != nil {
		return nil, err
	}
	return &player, nil
}
