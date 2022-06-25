package router

import (
	"002/appctx"

	"github.com/gin-gonic/gin"

	permissiongin "002/modules/permission/transport"
)

func MainRoute(g *gin.RouterGroup, appCtx appctx.AppContext) {
	permission := g.Group("/permission")
	{
		permission.POST("/insert", permissiongin.InsertPermission(appCtx))
		permission.GET("", permissiongin.SelectAll(appCtx))
		permission.DELETE("/delete/:id", permissiongin.Delete(appCtx))
	}
}
