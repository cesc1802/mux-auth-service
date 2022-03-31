package userstorage

import (
	"auth-service/common"
	"auth-service/modules/user_v1/model"
	"context"
)

func (s *sqlStore) Create(ctx context.Context, data *model.UserModel) error {
	tx := s.db.Begin()

	if err := tx.Create(data).Error; err != nil {
		tx.Rollback()
		return common.ErrDB(err)
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return common.ErrDB(err)
	}

	return nil
}
