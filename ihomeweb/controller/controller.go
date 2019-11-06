package controller

import (
	contextSys "context"
	"encoding/json"
	"fmt"
	"github.com/afocus/captcha"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/service/grpc"
	"github.com/micro/go-plugins/registry/consul"
	"image/png"
	getArea "microProject/ihomeweb/getAreaproto"
	getImage "microProject/ihomeweb/getImagecdproto"
	getSmscd "microProject/ihomeweb/getSmscdproto"
	"microProject/ihomeweb/utils"
	"regexp"
)

func GetArea(ctx *gin.Context) {
	//处理数据
	//链接服务
	reg := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"192.168.137.130:8500",
		}
	})
	grpcService := grpc.NewService(
		micro.Registry(reg),
	)
	client := getArea.NewGetAreaService("go.micro.srv.getArea", grpcService.Client())
	resp, err := client.Call(contextSys.TODO(), &getArea.Request{})
	if err != nil {
		ctx.JSON(404, nil)
		return
	}
	ctx.JSON(200, resp)
}
func GetSession(ctx *gin.Context) {
	s:=sessions.Default(ctx)
	s.Set("name","mrhu")
	s.Save()
	name:=s.Get("name")
	fmt.Println(name)

	resp := make(map[string]interface{})
	resp["errno"] = utils.RECODE_SESSIONERR
	resp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)



	ctx.JSON(200, resp)


}
func GetImageCd(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	grpcService := utils.GetGrpcService()
	client := getImage.NewGetImagecdService("go.micro.srv.getImagecd", grpcService.Client())
	resp, err := client.Call(contextSys.TODO(), &getImage.Request{Uuid: uuid})
	if err != nil {
		resp.Errno = utils.RECODE_DATAERR
		resp.Errmsg = utils.RecodeText(utils.RECODE_DATAERR)
		ctx.JSON(200, resp)
		return
	}
	var image captcha.Image
	json.Unmarshal(resp.Img, &image)
	png.Encode(ctx.Writer, image)
}
func GetSmscd(ctx *gin.Context) {
	resp := make(map[string]interface{})
	phone := ctx.Param("mobile")
	text := ctx.Query("text")
	id := ctx.Query("id")
	if phone == "" || text == "" || id == "" {
		resp["errno"] = utils.RECODE_DATAERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_DATAERR)
		ctx.JSON(200, resp)
		return
	}
	reg := regexp.MustCompile(`^1(3|4|5|6|7|8|9)\d{9}$`)
	result := reg.MatchString(phone)
	if result == false {
		resp["errno"] = utils.RECODE_MOBILEERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_MOBILEERR)
		ctx.JSON(200, resp)
		return
	}

	config := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"192.168.137.130:8500",
		}
	})
	grpcService := grpc.NewService(
		micro.Registry(config),
	)
	client := getSmscd.NewGetSmscdService("go.micro.srv.getSmscd", grpcService.Client())
	res, err := client.Call(contextSys.TODO(), &getSmscd.Request{
		Phone: phone,
		Text:  text,
		Uuid:  id,
	})
	if err != nil {
		resp["errno"] = utils.RECODE_DATAERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_DATAERR)
		ctx.JSON(200, resp)
		return
	}
	ctx.JSON(200,res)
}
