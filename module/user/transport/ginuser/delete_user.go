package ginuser

import (
	"net/http"
	"nevad/common"
	"nevad/component/appctx"
	userbiz "nevad/module/user/biz"
	userstore "nevad/module/user/store"

	"github.com/gin-gonic/gin"
)

func DeleteUser(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetMaiDBConnection()

		//id, err := strconv.Atoi(c.Param("id"))

		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := userstore.NewSQLStore(db)
		biz := userbiz.NewDeleteUserBiz(store)

		if err := biz.DeleteUser(c.Request.Context(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}