package transport

import (
	"002/appctx"
	"002/modules/permission/biz"
	"002/modules/permission/reponi"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SelectAll(appctx appctx.AppContext) func(g *gin.Context) {
	return func(g *gin.Context) {
		repo := reponi.NewRepoPermission(appctx.GetMainDBConnection())
		fmt.Println(repo)
		biz := biz.NewBizPermission(repo)
		all, err := biz.SelectAll()
		if err != nil {
			g.JSON(http.StatusBadRequest, err)
			return
		}

		g.JSON(http.StatusOK, all)
	}
}
