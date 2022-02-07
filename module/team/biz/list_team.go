package teambiz

import (
	"context"
	"nevad/common"
	teammodel "nevad/module/team/model"
)

type ListTeamStore interface {
	ListDataWithCondition(
		context context.Context,
		filter *teammodel.Filter,
		paging *common.Paging,
		morekeys ...string,
	) ([]teammodel.Team, error)
}

type listTeamBiz struct {
	store ListTeamStore
}

func NewListTeamBiz(store ListTeamStore) *listTeamBiz {
	return &listTeamBiz{
		store: store,
	}
}

func (biz listTeamBiz) ListTeam(
	context context.Context,
	filter *teammodel.Filter,
	paging *common.Paging,
) ([]teammodel.Team, error) {

	result, err := biz.store.ListDataWithCondition(context, filter, paging)

	if err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
