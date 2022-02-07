package profilebiz

import (
	"context"
	"nevad/common"

	profilemodel "nevad/module/profile/model"
)

// interface create profile for business
type CreateProfileStore interface {

	// method get profile by condiftion
	FindProfile(
		ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string,
	) (*profilemodel.Profile, error)

	// method get a proxy by condition
	FindProxy(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*common.SimpleProxy, error)

	// method create a profile with parameter
	CreateProfile(
		ctx context.Context,
		data *profilemodel.CreateProfile,
	) error
}

// struc profile business
type createProfileBiz struct {
	store CreateProfileStore
}

func NewCreateProfileBiz(store CreateProfileStore) *createProfileBiz {
	return &createProfileBiz{
		store: store,
	}
}

func (biz *createProfileBiz) CreateProfile(ctx context.Context, data *profilemodel.CreateProfile) error {

	profile, _ := biz.store.FindProfile(ctx, map[string]interface{}{
		"id": data.Id,
	})

	if profile != nil {
		return profilemodel.ErrProfileExisted
	}

	// check proxy is exist
	if proxy_id := data.ProxyId; proxy_id != nil {
		proxy, _ := biz.store.FindProxy(ctx, map[string]interface{}{
			"id": data.ProxyId,
		})
		if proxy == nil {
			return profilemodel.ErrProxyNotExist
		}
	}

	if err := biz.store.CreateProfile(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(profilemodel.EntityName, err)
	}

	return nil
}
