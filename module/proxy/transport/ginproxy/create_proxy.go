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

func CreateProxy(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetMaiDBConnection()
		var data proxymodel.Proxy

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := proxystore.NewSQLStore(db)
		biz := proxybiz.NewCreateBusiness(store)

		if err := biz.CreateProxy(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask()

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
