package ginuser

import (
	"net/http"
	"nevad/common"
	"nevad/component/appctx"
	userbiz "nevad/module/user/biz"
	userstore "nevad/module/user/store"

	"github.com/gin-gonic/gin"
)

func GetUser(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		db := ctx.GetMaiDBConnection()

		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := userstore.NewSQLStore(db)
		biz := userbiz.GetUserStore(store)

		data, err := biz.FindDataWithCondition(c.Request.Context(), map[string]interface{}{"id": uid.GetLocalID()})

		if err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}