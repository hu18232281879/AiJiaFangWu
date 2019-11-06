package main

import (
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"microProject/ihomeweb/controller"
	"microProject/ihomeweb/model"
	"microProject/ihomeweb/utils"
	"time"
)

func IsLogin() gin.HandlerFunc {
	return func(context *gin.Context) {
		t := time.Now()
		context.Next()
		dua := time.Since(t)
		fmt.Println("这个事务消耗了时间为,", dua)
	}
}

func main() {
	engine := gin.Default()
	store,err:=sessions.NewRedisStore(10,"tcp",utils.Redis_Address+":"+utils.Redis_Port,"",[]byte("session"))
	if err!=nil{
		fmt.Println("session初始化失败")
		return
	}
	store.Options(sessions.Options{MaxAge:0})
	engine.Use(sessions.Sessions("Mysession",store))

	model.InitDb()
	engine.Static("/home", "view")

	//REST概念	路由风格
	r1 := engine.Group("api/v1.0")
	{
		//r1.Use(IsLogin())
		r1.GET("/areas", controller.GetArea)
		r1.GET("/session", controller.GetSession)
		r1.GET("/imagecode/:uuid", controller.GetImageCd)
		r1.GET("/smscode/:mobile", controller.GetSmscd)
	
	}



	engine.Run(":8080")

}
