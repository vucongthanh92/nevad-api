package teambiz

import (
	"context"
	"nevad/common"
	teammodel "nevad/module/team/model"
)

type DeleteTeamStore interface {
	FindDataWithCondition(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string) (*teammodel.Team, error)
	Delete(context context.Context, id int) error
}

type deleteTeamBiz struct {
	store DeleteTeamStore
}

func NewDeleteTeamBiz(store DeleteTeamStore) *deleteTeamBiz {
	return &deleteTeamBiz{store: store}
}

func (biz deleteTeamBiz) DeleteTeam(context context.Context, id int) error {

	_, err := biz.store.FindDataWithCondition(context, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrEntityNotFound(teammodel.EntityName, err)
	}

	if err := biz.store.Delete(context, id); err != nil {
		return common.ErrCannotDeleteEntity(teammodel.EntityName, err)
	}

	return nil
}
