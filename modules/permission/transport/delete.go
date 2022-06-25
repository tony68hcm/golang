package transport

import (
	"002/appctx"
	"002/modules/permission/biz"
	"002/modules/permission/model"
	"002/modules/permission/reponi"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Delete(ctx appctx.AppContext) func(g *gin.Context) {
	return func(g *gin.Context) {

		id, err := strconv.ParseInt(g.Param("id"), 10, 64)
		if err != nil {
			g.JSON(http.StatusBadRequest, model.Response{
				ReasonCode:    400,
				ReasonMessage: "Please Input Number Id",
				Status:        "FAIL",
			})
			return
		}

		//delete
		repo := reponi.NewRepoPermission(ctx.GetMainDBConnection())
		biz := biz.NewBizPermission(repo)
		err = biz.DeleteId(id)
		if err != nil {
			g.JSON(http.StatusBadRequest, model.Response{
				ReasonCode:    400,
				ReasonMessage: err.Error(),
				Status:        "FAIL",
			})
			return
		}

		g.JSON(http.StatusOK, model.Response{
			ReasonCode:    200,
			ReasonMessage: "Delete Success",
			Status:        "OK",
		})
	}
}
