package profilebiz

import (
	"context"
	"nevad/common"

	profilemodel "nevad/module/profile/model"
)

type UpdateProfileStore interface {

	// method get profile by condiftion
	FindProfile(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*profilemodel.Profile, error)

	// method get a proxy by condition
	FindProxy(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*common.SimpleProxy, error)

	// method update a profile with parameter
	Update(
		ctx context.Context,
		id int,
		data *profilemodel.UpdateProfile,
	) error
}

type updateProfileBiz struct {
	store UpdateProfileStore
}

func NewUpdateProfile(store UpdateProfileStore) *updateProfileBiz {
	return &updateProfileBiz{
		store: store,
	}
}

func (biz *updateProfileBiz) UpdateProfile(ctx context.Context, id int, data *profilemodel.UpdateProfile) error {

	_, err := biz.store.FindProfile(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return common.ErrEntityNotFound(profilemodel.EntityName, err)
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

	if err := biz.store.Update(ctx, id, data); err != nil {
		return common.ErrCannotUpdateEntity(profilemodel.EntityName, err)
	}

	data.Id = id

	return nil

}
