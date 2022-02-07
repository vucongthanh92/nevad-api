package ginuser

import (
	"fmt"
	"net/http"
	"nevad/common"
	"nevad/component/appctx"
	"nevad/component/hasher"
	userbiz "nevad/module/user/biz"
	usermodel "nevad/module/user/model"
	userstore "nevad/module/user/store"

	"github.com/gin-gonic/gin"
)

func UpdateUser(ctx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := ctx.GetMaiDBConnection()

		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var data usermodel.UserUpdate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := userstore.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()
		biz := userbiz.NewUpdateUserBiz(store, md5)

		fmt.Println("c.Request.Context()", c.Request.Context())

		if err := biz.UpdateUser(c.Request.Context(), int(uid.GetLocalID()), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}