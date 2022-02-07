package teamstore

import (
	"context"
	"nevad/common"
	teammodel "nevad/module/team/model"
)

func (s *sqlStore) Create(context context.Context, data *teammodel.TeamCreate) error {
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
