package proxystore

import (
	"context"
	"nevad/common"
	proxymodel "nevad/module/proxy/model"
)

func (s *sqlStore) Update(ctx context.Context, id int, data *proxymodel.Proxy) error {
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
