package userbiz

import (
	"context"
	"nevad/common"
	usermodel "nevad/module/user/model"
)

type UpdateUserStore interface {
	FindDataWithCondition(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*usermodel.User, error)
	Update(ctx context.Context, id int, data *usermodel.UserUpdate) error
}

type updateUserBiz struct {
	store UpdateUserStore
	hasher Hasher
}

func NewUpdateUserBiz(store UpdateUserStore, hasher Hasher) *updateUserBiz {
	return &updateUserBiz{
		store: store,
		hasher: hasher,
	}
}

func (biz updateUserBiz) UpdateUser(context context.Context, id int, data *usermodel.UserUpdate) error {
	user, err := biz.store.FindDataWithCondition(context, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrEntityNotFound(usermodel.EntityName, err)
	}

	if len(*data.Password) > 0 {
		newPassword := biz.hasher.Hash(*data.Password + user.Salt)
		data.Password = &newPassword
	}

	if err := biz.store.Update(context, id, data); err != nil {
		return common.ErrCannotUpdateEntity(usermodel.EntityName, err)
	}

	return nil
}