package userstore

import (
	"context"
	"nevad/common"
	usermodel "nevad/module/user/model"
)

func (s *sqlStore) ListDataWithCondition(
	context context.Context,
	filter *usermodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]usermodel.User, error) {
	var result []usermodel.User

	var sortField, sortOrder string = "id", "desc"

	db := s.db.Table(usermodel.User{}.TableName())

	if f := filter; f != nil {
		if f.OwnerId > 0 {
			db = db.Where("owner_id = ?", f.OwnerId)
		}

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
		last.Mask(false)
		paging.NextCursor = last.FakeId.String()
	}

	return result, nil
}