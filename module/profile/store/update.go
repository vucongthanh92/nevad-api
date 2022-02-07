package profilestore

import (
	"context"
	"nevad/common"

	profilemodel "nevad/module/profile/model"
)

func (s *sqlStore) Update(ctx context.Context, id int, data *profilemodel.UpdateProfile) error {

	db := s.db.Begin()

	if err := db.Table(data.TableName()).Select("profile_name", "user_agent", "proxy_id").
		Where("id = ?", id).Updates(data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil

}
