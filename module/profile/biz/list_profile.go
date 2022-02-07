package profilebiz

import (
	"context"
	"nevad/common"

	profilemodel "nevad/module/profile/model"
	proxymodel "nevad/module/proxy/model"
)

type ListProfileStore interface {
	ListDataWithCondition(
		ctx context.Context,
		filter *profilemodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]profilemodel.Profile, error)
}

type listProfileBiz struct {
	store ListProfileStore
}

func NewListProfileBiz(store ListProfileStore) *listProfileBiz {
	return &listProfileBiz{store: store}
}

func (biz *listProfileBiz) ListProfile(
	ctx context.Context,
	filter *profilemodel.Filter,
	paging *common.Paging,
) ([]profilemodel.Profile, error) {

	var (
		relateKey = proxymodel.EntityName
	)

	result, err := biz.store.ListDataWithCondition(ctx, filter, paging, relateKey)

	if err != nil {
		return nil, common.ErrCannotListEntity(profilemodel.EntityName, err)
	}

	return result, nil

}
