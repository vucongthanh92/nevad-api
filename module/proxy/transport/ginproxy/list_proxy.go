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

func ListProxy(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetMaiDBConnection()

		var pagingData common.Paging

		if err := c.ShouldBind(&pagingData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		pagingData.Fulfill()

		var filter proxymodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		// filter.Status = []int{1}

		store := proxystore.NewSQLStore(db)
		biz := proxybiz.NewListProxyBiz(store)

		result, err := biz.ListProxy(c.Request.Context(), &filter, &pagingData)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask()
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, pagingData, filter))
	}
}
