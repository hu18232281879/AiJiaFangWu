package main

import (
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"microProject/ihomeweb/controller"
	"microProject/ihomeweb/model"
	"microProject/ihomeweb/utils"
)

func IsLogin() gin.HandlerFunc {
	return func(ctx *gin.Context){
		s:=sessions.Default(ctx)
		userName:=s.Get("userName")
		if userName==nil{
			resp:=make(map[string]interface{})
			resp["errno"]=utils.RECODE_SESSIONERR
			resp["errmsg"]=utils.RecodeText(utils.RECODE_SESSIONERR)
			ctx.JSON(200,resp)
		}else {
			ctx.Next()
		}
	}
}

func main() {
	engine := gin.Default()
	store,err:=sessions.NewRedisStore(10,"tcp",utils.Redis_Address+":"+utils.Redis_Port,"",[]byte("session"))
	if err!=nil{
		fmt.Println("session初始化失败")
		return
	}
	store.Options(sessions.Options{MaxAge:60*60*24})
	engine.Use(sessions.Sessions("Mysession",store))
	model.InitDb()
	engine.Static("/home", "view")
	//REST概念	路由风格
	r1 := engine.Group("api/v1.0")
	{
		r1.GET("/areas", controller.GetArea)
		r1.GET("/session", controller.GetSession)
		r1.GET("/imagecode/:uuid", controller.GetImageCd)
		r1.GET("/smscode/:mobile", controller.GetSmscd)
		r1.POST("/users",controller.PostRet)
		r1.DELETE("/session",controller.DeleteSession)
		r1.POST("/sessions",controller.PostLogin)
		r1.Use(IsLogin())
		r1.GET("/user",controller.GetUserInfo)

	}
	
	

	engine.Run(":8081")

}
