package ginproxy

import (
	"net/http"
	"nevad/common"
	"nevad/component/appctx"
	proxybiz "nevad/module/proxy/biz"
	proxymodel "nevad/module/proxy/model"
	proxystore "nevad/module/proxy/store"

	"github.com/gin-gonic/gin"
)

func UpdateProxy(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := ctx.GetMaiDBConnection()

		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var data proxymodel.Proxy

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := proxystore.NewSQLStore(db)
		biz := proxybiz.NewUpdateProxyBiz(store)

		if err := biz.UpdateProxy(c.Request.Context(), int(uid.GetLocalID()), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
