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

func CreateProfile(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var (
			db   = appCtx.GetMaiDBConnection()
			data profilemodel.CreateProfile
		)

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		// convert fakeProxy to get proxyId
		if *data.FakeProxyId != "" {
			data.GetProxyIdFromFakeProxy()
		}

		store := profilestore.NewSQLStore(db)

		biz := profilebiz.NewCreateProfileBiz(store)

		if err := biz.CreateProfile(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask()

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))

	}
}
