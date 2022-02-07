package ginteam

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nevad/common"
	"nevad/component/appctx"
	teambiz "nevad/module/team/biz"
	teammodel "nevad/module/team/model"
	teamstore "nevad/module/team/store"
)

func ListTeam(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := ctx.GetMaiDBConnection()

		var pagingData common.Paging

		if err := c.ShouldBind(&pagingData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		pagingData.Fulfill()

		var filter teammodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := teamstore.NewSQLStore(db)
		biz := teambiz.ListTeamStore(store)

		result, err := biz.ListDataWithCondition(c.Request.Context(), &filter, &pagingData)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, pagingData, filter))
	}
}
