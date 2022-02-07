package teamstore

import (
	"context"
	"nevad/common"
	teammodel "nevad/module/team/model"
)

func (s *sqlStore) ListDataWithCondition(
	context context.Context,
	filter *teammodel.Filter,
	paging *common.Paging,
	morekeys ...string,
) ([]teammodel.Team, error) {
	var result []teammodel.Team

	db := s.db.Table(teammodel.Team{}.TableName()).Where("status in (1)")

	if f := filter; f != nil {
		if f.OwnerId > 0 {
			db = db.Where("owner_id = ?", f.OwnerId)
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

		db = db.Where("id < ? ", uid.GetLocalID())
	} else {
		offset := (paging.Page - 1) * paging.Limit
		db = db.Offset(offset)
	}

	if err := db.
		Limit(paging.Limit).
		Order("id desc").
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
