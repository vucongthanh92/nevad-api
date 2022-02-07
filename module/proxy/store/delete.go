package proxystore

import (
	"context"
	"nevad/common"
	proxymodel "nevad/module/proxy/model"
)

func (s *sqlStore) Delete(context context.Context, id int) error {
	if err := s.db.Table(proxymodel.Proxy{}.TableName()).
		Where("id = ?", id).Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}