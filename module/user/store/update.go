package userstore

import (
	"context"
	"nevad/common"
	usermodel "nevad/module/user/model"
)

func (s *sqlStore) Update(ctx context.Context, id int, data *usermodel.UserUpdate) error {
	db := s.db.Begin()

	if err := db.Table(data.TableName()).
		Where("id = ?", id).
		Updates(data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}