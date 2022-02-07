package ginprofile

import (
	"net/http"
	"nevad/common"
	"nevad/component/appctx"
	profilebiz "nevad/module/profile/biz"
	profilestore "nevad/module/profile/store"

	"github.com/gin-gonic/gin"
)

func DeleteProfile(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {

		db := appCtx.GetMaiDBConnection()

		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := profilestore.NewSQLStore(db)
		biz := profilebiz.NewDeleteProfileBiz(store)

		if err := biz.DeleteProfile(c.Request.Context(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))

	}
}
