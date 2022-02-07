package ginteam

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nevad/common"
	"nevad/component/appctx"
	teambiz "nevad/module/team/biz"
	teamstore "nevad/module/team/store"
)

func DeleteTeam(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := ctx.GetMaiDBConnection()

		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := teamstore.NewSQLStore(db)
		biz := teambiz.NewDeleteTeamBiz(store)

		if err := biz.DeleteTeam(c.Request.Context(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
