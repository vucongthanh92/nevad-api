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

func UpdateTeam(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := ctx.GetMaiDBConnection()

		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var data teammodel.TeamUpdate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := teamstore.NewSQLStore(db)
		biz := teambiz.NewUpdateTeamBiz(store)

		if err := biz.UpdateTeam(c.Request.Context(), int(uid.GetLocalID()), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
