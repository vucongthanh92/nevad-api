package ginuser

import (
	"net/http"
	"nevad/common"
	"nevad/component/appctx"
	userbiz "nevad/module/user/biz"
	usermodel "nevad/module/user/model"
	userstore "nevad/module/user/store"

	"github.com/gin-gonic/gin"
)

func ListUser(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetMaiDBConnection()

		var pagingData common.Paging

		if err := c.ShouldBind(&pagingData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		pagingData.Fulfill()

		var filter usermodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		// filter.Status = []int{1}

		store := userstore.NewSQLStore(db)
		biz := userbiz.NewListUserBiz(store)

		result, err := biz.ListRestaurant(c.Request.Context(), &filter, &pagingData)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, pagingData, filter))
	}
}
