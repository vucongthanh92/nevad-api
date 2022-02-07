package profilestore

import (
	"context"
	"fmt"
	"nevad/common"

	"gorm.io/gorm"

	profilemodel "nevad/module/profile/model"
)

func (s *sqlStore) FindProfile(
	ctx context.Context,
	conditions map[string]interface{},
	moreKeys ...string,
) (*profilemodel.Profile, error) {

	var (
		db      = s.db.Table(profilemodel.Profile{}.TableName())
		profile profilemodel.Profile
	)

	// create a loop to get string into array moreKeys.
	// it will be add into preload to get mutil table
	if len(moreKeys) > 0 {
		for i := range moreKeys {
			db = db.Preload(moreKeys[i])
		}
	}

	if err := db.Where(conditions).First(&profile).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}

		return nil, common.ErrDB(err)
	}

	//=====================================================
	tx := s.db.Begin()

	newProfile := profilemodel.Profile{
		ProfileName: "francis02",
		UserAgent:   "hello world",
		ProxyId:     1,
		Status:      true,
	}

	err := tx.Create(&newProfile).Error
	if err != nil {
		tx.Rollback()
		fmt.Println("err:", err)
	}

	newProfile.ProfileName = "francis 03"

	err = tx.Table(profile.TableName()).
		Where("id = ?", profile.Id).
		Update("profile_name", newProfile.ProfileName).Error
	if err != nil {
		tx.Rollback()
		fmt.Println("err:", err)
	}
	tx.Commit()
	//=====================================================

	return &profile, nil
}

// FindProxy func will get a proxy by ID
func (s *sqlStore) FindProxy(
	ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*common.SimpleProxy, error) {

	var proxy common.SimpleProxy

	if err := s.db.Where(condition).First(&proxy).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}

	return &proxy, nil
}
