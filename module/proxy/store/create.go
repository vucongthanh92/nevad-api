package proxystore

import (
	"context"
	"nevad/common"
	proxymodel "nevad/module/proxy/model"
)

func (s *sqlStore) CreateProxy(ctx context.Context, data *proxymodel.Proxy) error {
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
