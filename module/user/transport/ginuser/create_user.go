package ginuser

import (
	"net/http"
	"nevad/common"
	"nevad/component/appctx"
	"nevad/component/hasher"
	userbiz "nevad/module/user/biz"
	usermodel "nevad/module/user/model"
	userstore "nevad/module/user/store"

	"github.com/gin-gonic/gin"
)

func Create(appCtx appctx.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetMaiDBConnection()
		var data usermodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := userstore.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()
		biz := userbiz.NewCreateBusiness(store, md5)

		if err := biz.Create(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}