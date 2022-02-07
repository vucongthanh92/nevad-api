package userbiz

import (
	"context"
	"nevad/common"
	usermodel "nevad/module/user/model"
)

type GetUserStore interface {
	FindDataWithCondition(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*usermodel.User, error)
}

type getUserBiz struct {
	store GetUserStore
}

func NewGetUserBiz(store GetUserStore) *getUserBiz {
	return &getUserBiz{
		store: store,
	}
}

func (biz getUserBiz) GetUser(
	context context.Context,
	condition map[string]interface{},
) (*usermodel.User, error) {
	result, err := biz.store.FindDataWithCondition(context, condition)

	if err != nil {
		return nil, common.ErrCannotGetEntity(usermodel.EntityName, err)
	}

	return result, nil
}