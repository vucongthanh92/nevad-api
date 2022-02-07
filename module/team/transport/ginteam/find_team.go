package ginteam

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nevad/common"
	"nevad/component/appctx"
	teambiz "nevad/module/team/biz"
	teamstore "nevad/module/team/store"
)

func FindTeam(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		db := ctx.GetMaiDBConnection()

		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := teamstore.NewSQLStore(db)
		biz := teambiz.FindTeamStore(store)

		data, err := biz.FindDataWithCondition(c.Request.Context(), map[string]interface{}{"id": uid.GetLocalID()})

		if err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
