package teambiz

import (
	"context"
	"nevad/common"
	teammodel "nevad/module/team/model"
)

type UpdateTeamStore interface {
	Update(ctx context.Context, id int, data *teammodel.TeamUpdate) error
}

type updateTeamBiz struct {
	store UpdateTeamStore
}

func NewUpdateTeamBiz(store UpdateTeamStore) *updateTeamBiz {
	return &updateTeamBiz{store: store}
}

func (biz updateTeamBiz) UpdateTeam(context context.Context, id int, data *teammodel.TeamUpdate) error {

	if err := biz.store.Update(context, id, data); err != nil {
		return common.ErrCannotUpdateEntity(teammodel.EntityName, err)
	}

	return nil
}