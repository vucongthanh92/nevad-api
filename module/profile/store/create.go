package profilestore

import (
	"context"
	"nevad/common"

	profilemodel "nevad/module/profile/model"
)

func (s *sqlStore) CreateProfile(ctx context.Context, data *profilemodel.CreateProfile) error {

	db := s.db.Begin()

	if err := db.Table(data.TableName()).Create(data).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil

}
