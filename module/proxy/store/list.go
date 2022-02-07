package proxystore

import (
	"context"
	"nevad/common"
	proxymodel "nevad/module/proxy/model"
)

func (s *sqlStore) ListDataWithCondition(
	context context.Context,
	filter *proxymodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]proxymodel.Proxy, error) {
	var result []proxymodel.Proxy

	var sortField, sortOrder string = "id", "desc"

	db := s.db.Table(proxymodel.Proxy{}.TableName())

	if f := filter; f != nil {

		if len(f.Status) > 0 {
			db = db.Where("status in (?)", f.Status)
		}

		if len(f.SortField) > 0 {
			sortField = f.SortField
		}

		if len(f.SortOrder) > 0 {
			sortOrder = f.SortOrder
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if v := paging.FakeCursor; v != "" {
		uid, err := common.FromBase58(v)

		if err != nil {
			return nil, common.ErrDB(err)
		}

		db = db.Where("id < ?", uid.GetLocalID())
	} else {
		offset := (paging.Page - 1) * paging.Limit
		db = db.Offset(offset)
	}

	if err := db.
		Limit(paging.Limit).
		Order(sortField + " " + sortOrder).
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if len(result) > 0 {
		last := result[len(result)-1]
		last.Mask()
		paging.NextCursor = last.FakeId.String()
	}

	return result, nil
}
