package proxybiz

import (
	"context"
	"nevad/common"
	proxymodel "nevad/module/proxy/model"
)

type DeleteProxyStore interface {
	FindDataWithCondition(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*proxymodel.Proxy, error)
	Delete(context context.Context, id int) error
}

type deleteProxyBiz struct {
	store DeleteProxyStore
}

func NewDeleteProxyBiz(store DeleteProxyStore) *deleteProxyBiz {
	return &deleteProxyBiz{store: store}
}

func (biz *deleteProxyBiz) DeleteProxy(context context.Context, id int) error {
	_, err := biz.store.FindDataWithCondition(context, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrEntityNotFound(proxymodel.EntityName, err)
	}

	if err := biz.store.Delete(context, id); err != nil {
		return common.ErrCannotDeleteEntity(proxymodel.EntityName, nil)
	}

	return nil
}