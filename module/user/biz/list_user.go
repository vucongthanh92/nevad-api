package userbiz

import (
	"context"
	"nevad/common"
	usermodel "nevad/module/user/model"
)

type ListUserStore interface {
	ListDataWithCondition(
		context context.Context,
		filter *usermodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]usermodel.User, error)
}

type listUserBiz struct {
	store ListUserStore
}

func NewListUserBiz(store ListUserStore) *listUserBiz {
	return &listUserBiz{store: store}
}

func (biz *listUserBiz) ListRestaurant(
	context context.Context,
	filter *usermodel.Filter,
	paging *common.Paging,
) ([]usermodel.User, error) {
	result, err := biz.store.ListDataWithCondition(context, filter, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(usermodel.EntityName, err)
	}

	return result, nil
}
