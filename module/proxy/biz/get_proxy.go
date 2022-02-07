package proxybiz

import (
	"context"
	"nevad/common"
	proxymodel "nevad/module/proxy/model"
)

type GetProxyStore interface {
	FindDataWithCondition(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*proxymodel.Proxy, error)
}

type getProxyBiz struct {
	store GetProxyStore
}

func NewGetProxyBiz(store GetProxyStore) *getProxyBiz {
	return &getProxyBiz{
		store: store,
	}
}

func (biz getProxyBiz) GetProxy(
	context context.Context,
	condition map[string]interface{},
) (*proxymodel.Proxy, error) {
	result, err := biz.store.FindDataWithCondition(context, condition)

	if err != nil {
		return nil, common.ErrCannotGetEntity(proxymodel.EntityName, err)
	}

	return result, nil
}
