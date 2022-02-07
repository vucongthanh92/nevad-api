package proxybiz

import (
	"context"
	"nevad/common"
	proxymodel "nevad/module/proxy/model"
)

type CreateStorage interface {
	FindDataWithCondition(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*proxymodel.Proxy, error)
	CreateProxy(ctx context.Context, data *proxymodel.Proxy) error
}

type createBusiness struct {
	createStorage CreateStorage
}

func NewCreateBusiness(createStorage CreateStorage) *createBusiness {
	return &createBusiness{
		createStorage: createStorage,
	}
}

func (business *createBusiness) CreateProxy(ctx context.Context, data *proxymodel.Proxy) error {
	proxy, _ := business.createStorage.FindDataWithCondition(ctx, map[string]interface{}{"ip": data.IP})

	if proxy != nil {
		return proxymodel.ErrIpExisted
	}

	if err := business.createStorage.CreateProxy(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(proxymodel.EntityName, err)
	}

	return nil
}
