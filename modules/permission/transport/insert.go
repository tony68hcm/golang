package transport

import (
	"002/appctx"
	"002/modules/permission/model"
	"net/http"

	"002/modules/permission/reponi"

	bizper "002/modules/permission/biz"

	"github.com/gin-gonic/gin"
)

func InsertPermission(appCtx appctx.AppContext) func(g *gin.Context) {
	return func(g *gin.Context) {
		req := model.Permission{}

		err := g.ShouldBindJSON(&req)
		if err != nil {
			g.JSON(http.StatusBadRequest, err)
			return
		}

		//insert vo table permission
		store := reponi.NewRepoPermission(appCtx.GetMainDBConnection())
		biz := bizper.NewBizPermission(store)
		err = biz.InsertPermission(req)
		if err != nil {
			g.JSON(http.StatusBadRequest, err)
			return
		}

		g.JSON(http.StatusOK, model.Response{
			ReasonMessage: "insert thành công",
			ReasonCode:    200,
			Status:        "OK",
		})

	}
}
