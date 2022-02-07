package teambiz

import (
	"context"
	"nevad/common"
	teammodel "nevad/module/team/model"
)

type CreateTeamStore interface {
	Create(context context.Context, data *teammodel.TeamCreate) error
}

type createTeamBiz struct {
	store CreateTeamStore
}

func NewCreateTeamBiz(store CreateTeamStore) *createTeamBiz {
	return &createTeamBiz{store: store}
}

func (biz createTeamBiz) CreateTeam(context context.Context, data *teammodel.TeamCreate) error {

	// VALIDATE
	if err := data.Validate(); err != nil {
		return common.ErrInvalidRequest(err)
	}

	if err := biz.store.Create(context, data); err != nil {
		return common.ErrCannotCreateEntity(teammodel.EntityName, err)
	}

	return nil
}
