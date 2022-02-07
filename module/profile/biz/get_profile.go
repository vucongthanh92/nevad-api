package profilebiz

import (
	"context"
	"nevad/common"

	profilemodel "nevad/module/profile/model"
	proxymodel "nevad/module/proxy/model"
)

type GetProfileStore interface {
	FindProfile(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*profilemodel.Profile, error)
}

type getProfileBiz struct {
	store GetProfileStore
}

func NewGetProfileBiz(store GetProfileStore) *getProfileBiz {
	return &getProfileBiz{
		store: store,
	}
}

func (biz *getProfileBiz) GetProfile(ctx context.Context, condition map[string]interface{}) (*profilemodel.Profile, error) {

	var (
		relateKey = proxymodel.EntityName
	)

	result, err := biz.store.FindProfile(ctx, condition, relateKey)

	if err != nil {
		return nil, common.ErrCannotGetEntity(profilemodel.EntityName, err)
	}

	return result, nil

}
