package proxybiz

import (
	"context"
	"nevad/common"
	proxymodel "nevad/module/proxy/model"
)

type UpdateProxyStore interface {
	FindDataWithCondition(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*proxymodel.Proxy, error)
	Update(ctx context.Context, id int, data *proxymodel.Proxy) error
}

type updateProxyBiz struct {
	store UpdateProxyStore
}

func NewUpdateProxyBiz(store UpdateProxyStore) *updateProxyBiz {
	return &updateProxyBiz{
		store: store,
	}
}

func (biz updateProxyBiz) UpdateProxy(context context.Context, id int, data *proxymodel.Proxy) error {
	_, err := biz.store.FindDataWithCondition(context, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrEntityNotFound(proxymodel.EntityName, err)
	}

	if err := biz.store.Update(context, id, data); err != nil {
		return common.ErrCannotUpdateEntity(proxymodel.EntityName, err)
	}

	return nil
}
