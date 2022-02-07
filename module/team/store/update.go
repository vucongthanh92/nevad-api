package teamstore

import (
	"context"
	"nevad/common"
	teammodel "nevad/module/team/model"
)

func (s *sqlStore) Update(ctx context.Context, id int, data *teammodel.TeamUpdate) error {
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
