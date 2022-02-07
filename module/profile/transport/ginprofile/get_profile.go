package ginprofile

import (
	"net/http"
	"nevad/common"
	"nevad/component/appctx"

	profilebiz "nevad/module/profile/biz"
	profilestore "nevad/module/profile/store"

	"github.com/gin-gonic/gin"
)

func GetProfile(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		db := ctx.GetMaiDBConnection()

		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := profilestore.NewSQLStore(db)
		biz := profilebiz.NewGetProfileBiz(store)

		data, err := biz.GetProfile(c.Request.Context(), map[string]interface{}{
			"id": uid.GetLocalID(),
		})

		if err != nil {
			panic(err)
		}

		data.Mask()

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
