package teambiz

import (
	"context"
	"nevad/common"
	teammodel "nevad/module/team/model"
)

type FindTeamStore interface {
	FindDataWithCondition(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*teammodel.Team, error)
}

type findTeamBiz struct {
	store FindTeamStore
}

func NewFindTeamBiz(store FindTeamStore) *findTeamBiz {
	return &findTeamBiz{
		store: store,
	}
}

func (biz findTeamBiz) FindTeam(
	context context.Context,
	condition map[string]interface{},
) (*teammodel.Team, error) {
	result, err := biz.store.FindDataWithCondition(context, condition)

	if err != nil {
		return nil, common.ErrCannotGetEntity(teammodel.EntityName, err)
	}

	return result, nil
}
