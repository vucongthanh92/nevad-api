package ginprofile

import (
	"net/http"
	"nevad/common"
	"nevad/component/appctx"

	profilebiz "nevad/module/profile/biz"
	profilemodel "nevad/module/profile/model"
	profilestore "nevad/module/profile/store"

	"github.com/gin-gonic/gin"
)

func ListProfile(appCtx appctx.AppContext) func(c *gin.Context) {

	return func(c *gin.Context) {

		var (
			db         = appCtx.GetMaiDBConnection()
			pagingData common.Paging
		)

		if err := c.ShouldBind(&pagingData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		pagingData.Fulfill()

		var filter profilemodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := profilestore.NewSQLStore(db)
		biz := profilebiz.NewListProfileBiz(store)

		result, err := biz.ListProfile(c.Request.Context(), &filter, &pagingData)

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask()
		}

		// return result with paging and filter in response
		c.JSON(http.StatusOK, common.NewSuccessResponse(result, pagingData, filter))
	}
}
