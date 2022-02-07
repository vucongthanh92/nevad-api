package ginuser

import (
	"net/http"
	"nevad/common"
	"nevad/component/appctx"

	"github.com/gin-gonic/gin"
)

func Profile(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		u := c.MustGet(common.CurrentUser)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(u))
	}
}