package profilestore

import (
	"context"
	"nevad/common"
	profilemodel "nevad/module/profile/model"
)

func (s *sqlStore) Delete(ctx context.Context, id int) error {

	if err := s.db.Table(profilemodel.Profile{}.TableName()).
		Where("id = ?", id).Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil

}
