package userstorage

import (
	"auth-service/common"
	"auth-service/modules/user_v1/model"
	"context"
)

func (s *sqlStore) UpdateUser(ctx context.Context, condition map[string]interface{}, data *model.UserModel) error {
	db := s.db.Table(data.TableName())

	if err := db.Where(condition).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
