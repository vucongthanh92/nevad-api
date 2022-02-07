package profilebiz

import (
	"context"
	"nevad/common"
	profilemodel "nevad/module/profile/model"
)

type DeleteProfileStore interface {
	FindProfile(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*profilemodel.Profile, error)
	Delete(ctx context.Context, id int) error
}

type deleteProfileBiz struct {
	store DeleteProfileStore
}

func NewDeleteProfileBiz(store DeleteProfileStore) *deleteProfileBiz {
	return &deleteProfileBiz{
		store: store,
	}
}

func (biz *deleteProfileBiz) DeleteProfile(ctx context.Context, id int) error {

	// check profile is exist
	_, err := biz.store.FindProfile(ctx, map[string]interface{}{
		"id": id,
	})

	if err != nil {
		return common.ErrEntityNotFound(profilemodel.EntityName, err)
	}

	if err := biz.store.Delete(ctx, id); err != nil {
		return common.ErrCannotDeleteEntity(profilemodel.EntityName, nil)
	}

	return nil

}
