package teamstore

import (
	"context"
	"gorm.io/gorm"
	"nevad/common"
	teammodel "nevad/module/team/model"
)

func (s *sqlStore) FindDataWithCondition(
	context context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*teammodel.Team, error) {

	var result teammodel.Team

	if err := s.db.Where(condition).Find(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}

		return nil, err
	}

	return &result, nil
}
