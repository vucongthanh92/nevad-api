package teamstore

import (
	"context"
	"nevad/common"
	teammodel "nevad/module/team/model"
)

func (s *sqlStore) Delete(context context.Context, id int) error {
	db := s.db.Begin()

	if err := db.Table(teammodel.Team{}.TableName()).
		Where("id = ?", id).
		Updates(map[string]interface{}{"status": 0}).Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	if err := db.Commit().Error; err != nil {
		db.Rollback()
		return common.ErrDB(err)
	}

	return nil
}
