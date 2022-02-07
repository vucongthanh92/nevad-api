package proxystore

import (
	"context"
	"nevad/common"
	proxymodel "nevad/module/proxy/model"

	"gorm.io/gorm"
)

func (s *sqlStore) FindProxy(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*proxymodel.Proxy, error) {
	db := s.db.Table(proxymodel.Proxy{}.TableName())

	for i := range moreInfo {
		db = db.Preload(moreInfo[i])
	}

	var proxy proxymodel.Proxy

	if err := db.Where(conditions).First(&proxy).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}

		return nil, common.ErrDB(err)
	}

	return &proxy, nil
}

func (s *sqlStore) FindDataWithCondition(
	context context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*proxymodel.Proxy, error) {
	var data proxymodel.Proxy

	if err := s.db.Where(condition).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &data, nil
}
