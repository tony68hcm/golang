package main

import (
	"002/appctx"
	"002/router"
	"fmt"

	"github.com/gin-gonic/gin"
	//"github.com/rs/cors/wrapper/gin"
)

func main() {
	fmt.Println("start")

	database, err := appctx.InitMysqlDB("sql6.freemysqlhosting.net", 3306, "sql6501022", "szkJpYD8aq", "sql6501022")
	if err != nil {
		panic(err)
	}
	appCtx := appctx.NewAppContext(database)
	fmt.Print(appCtx)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/v1")

	router.MainRoute(v1, appCtx)
	/*
		permission := v1.Group("/permission")
		{
			permission.POST("/insert", permissiongin.InsertPermission(appCtx))
		}
	*/

	if err := r.Run(); err != nil {
		fmt.Println("error start")
	}
}
