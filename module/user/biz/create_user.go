package userbiz

import (
	"context"
	"nevad/common"
	usermodel "nevad/module/user/model"
)

type CreateStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
}

type Hasher interface {
	Hash(data string) string
}

type createBusiness struct {
	createStorage CreateStorage
	hasher          Hasher
}

func NewCreateBusiness(createStorage CreateStorage, hasher Hasher) *createBusiness {
	return &createBusiness{
		createStorage: createStorage,
		hasher:          hasher,
	}
}

func (business *createBusiness) Create(ctx context.Context, data *usermodel.UserCreate) error {
	user, _ := business.createStorage.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if user != nil {
		//if user.Status == 0 {
		//	return error user has been disable
		//}

		return usermodel.ErrEmailExisted
	}

	salt := common.GenSalt(50)

	data.Password = business.hasher.Hash(data.Password + salt)
	data.Salt = salt
	//data.Status = 1

	if err := business.createStorage.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}

	return nil
}