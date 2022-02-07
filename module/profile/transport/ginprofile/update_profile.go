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

func UpdateProfile(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var (
			db   = ctx.GetMaiDBConnection()
			data profilemodel.UpdateProfile
		)

		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		// convert fakeProxy to get proxyId
		if *data.FakeProxyId != "" {
			data.GetProxyIdFromFakeProxy()
		}

		store := profilestore.NewSQLStore(db)
		biz := profilebiz.NewUpdateProfile(store)

		if err := biz.UpdateProfile(c.Request.Context(), int(uid.GetLocalID()), &data); err != nil {
			panic(err)
		}

		data.Mask()

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))

	}
}
