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

func CreateTeam(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := ctx.GetMaiDBConnection()

		var data teammodel.TeamCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := teamstore.NewSQLStore(db)
		biz := teambiz.NewCreateTeamBiz(store)

		if err := biz.CreateTeam(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId))
	}
}
