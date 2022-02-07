package userstore

import (
	"context"
	"nevad/common"
	usermodel "nevad/module/user/model"
)

func (s *sqlStore) Delete(context context.Context, id int) error {
	if err := s.db.Table(usermodel.User{}.TableName()).
		Where("id = ?", id).
		Updates(map[string]interface{}{"status": 0}).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}