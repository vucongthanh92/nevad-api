package ginproxy

import (
	"net/http"
	"nevad/common"
	"nevad/component/appctx"
	proxybiz "nevad/module/proxy/biz"
	proxystore "nevad/module/proxy/store"

	"github.com/gin-gonic/gin"
)

func DeleteProxy(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetMaiDBConnection()

		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := proxystore.NewSQLStore(db)
		biz := proxybiz.NewDeleteProxyBiz(store)

		if err := biz.DeleteProxy(c.Request.Context(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}