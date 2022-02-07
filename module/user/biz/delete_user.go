package userbiz

import (
	"context"
	"nevad/common"
	usermodel "nevad/module/user/model"
)

type DeleteUserStore interface {
	FindDataWithCondition(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*usermodel.User, error)
	Delete(context context.Context, id int) error
}

type deleteUserBiz struct {
	store DeleteUserStore
}

func NewDeleteUserBiz(store DeleteUserStore) *deleteUserBiz {
	return &deleteUserBiz{store: store}
}

func (biz *deleteUserBiz) DeleteUser(context context.Context, id int) error {
	oldData, err := biz.store.FindDataWithCondition(context, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrEntityNotFound(usermodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(usermodel.EntityName, nil)
	}

	if err := biz.store.Delete(context, id); err != nil {
		return common.ErrCannotDeleteEntity(usermodel.EntityName, nil)
	}

	return nil
}