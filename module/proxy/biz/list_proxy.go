package proxybiz

import (
	"context"
	"nevad/common"
	proxymodel "nevad/module/proxy/model"
)

type ListProxyStore interface {
	ListDataWithCondition(
		context context.Context,
		filter *proxymodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]proxymodel.Proxy, error)
}

type listProxyBiz struct {
	store ListProxyStore
}

func NewListProxyBiz(store ListProxyStore) *listProxyBiz {
	return &listProxyBiz{store: store}
}

func (biz *listProxyBiz) ListProxy(
	context context.Context,
	filter *proxymodel.Filter,
	paging *common.Paging,
) ([]proxymodel.Proxy, error) {
	result, err := biz.store.ListDataWithCondition(context, filter, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(proxymodel.EntityName, err)
	}

	return result, nil
}
