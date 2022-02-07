package profilestore

import (
	"context"
	"nevad/common"

	profilemodel "nevad/module/profile/model"
)

func (s *sqlStore) ListDataWithCondition(
	ctx context.Context,
	filter *profilemodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]profilemodel.Profile, error) {

	// initialize some variable will be used whole func.
	// result contain the data that get from table and it is array profile
	var (
		result               []profilemodel.Profile
		sortField, sortOrder string = "id", "desc"
	)

	// choose the table that you want to process with it.
	// The name of table will will be get via func TableName
	db := s.db.Table(profilemodel.Profile{}.TableName())

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

	// count total row get into database
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	// create a loop to get string into array moreKeys.
	// it will be add into preload to get mutil table
	if len(moreKeys) > 0 {
		for i := range moreKeys {
			db = db.Preload(moreKeys[i])
		}
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

	return result, nil

}
