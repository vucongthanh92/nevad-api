package ginproxy

import (
	"net/http"
	"nevad/common"
	"nevad/component/appctx"
	proxybiz "nevad/module/proxy/biz"
	proxystore "nevad/module/proxy/store"

	"github.com/gin-gonic/gin"
)

func GetProxy(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		db := ctx.GetMaiDBConnection()

		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := proxystore.NewSQLStore(db)
		biz := proxybiz.GetProxyStore(store)

		data, err := biz.FindDataWithCondition(c.Request.Context(), map[string]interface{}{"id": uid.GetLocalID()})

		if err != nil {
			panic(err)
		}

		data.Mask()

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
