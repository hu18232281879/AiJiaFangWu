package controller

import (
	contextSys "context"
	"encoding/json"
	"fmt"

	//"fmt"
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
	postLogin "microProject/ihomeweb/postLoginproto"
	postRet "microProject/ihomeweb/postRetproto"
	getUserInfo "microProject/ihomeweb/getUserInfo"
	"microProject/ihomeweb/utils"
	"regexp"
)

type User struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
	Sms_code string `json:"sms_code"`
}
type LogUser struct {
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
}
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
	resp := make(map[string]interface{})
	tempMap := make(map[string]string)
	s := sessions.Default(ctx)
	userName := s.Get("userName")
	if userName == nil {
		resp["errno"] = utils.RECODE_SESSIONERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)
	} else {
		resp["errno"] = utils.RECODE_OK
		resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
		tempMap["name"] = userName.(string)
		resp["data"] = tempMap
	}

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
	ctx.JSON(200, res)
}
func PostRet(ctx *gin.Context) {
	result := make(map[string]interface{})
	var user User
	err := ctx.Bind(&user)
	if err != nil {
		result["errno"] = utils.RECODE_DATAERR
		result["errmsg"] = utils.RecodeText(utils.RECODE_DATAERR)
		ctx.JSON(200, result)
		return
	}
	grpcService := utils.GetGrpcService()
	client := postRet.NewPostRetService("go.micro.srv.postRet", grpcService.Client())
	req := postRet.Request{
		Mobile:   user.Mobile,
		Password: user.Password,
		SmsCode:  user.Sms_code,
	}
	resp, err := client.Call(contextSys.TODO(), &req)
	if err != nil {
		result["errno"] = utils.RECODE_DATAERR
		result["errmsg"] = utils.RecodeText(utils.RECODE_DATAERR)
		ctx.JSON(200, result)
		return
	}
	s := sessions.Default(ctx)
	s.Set("userName", user.Mobile)
	err = s.Save()
	if err != nil {
		resp.Errno = utils.RECODE_SESSIONERR
		resp.Errmsg = utils.RecodeText(utils.RECODE_SESSIONERR)
	}

	ctx.JSON(200, resp)

}
func DeleteSession(ctx *gin.Context) {
	s := sessions.Default(ctx)
	s.Delete("userName")
	result := make(map[string]string)
	err := s.Save()
	if err != nil {
		result["errno"] = utils.RECODE_DATAERR
		result["errmsg"] = utils.RecodeText(utils.RECODE_DATAERR)
	} else {
		result["errno"] = utils.RECODE_OK
		result["errmsg"] = utils.RecodeText(utils.RECODE_OK)
	}
	ctx.JSON(200, result)
}
func PostLogin(ctx *gin.Context) {
	errResult := make(map[string]string)
	var logUser LogUser
	err := ctx.BindJSON(&logUser)
	if logUser.Mobile == "" || logUser.Password == "" {
		errResult["errno"] = utils.RECODE_DATAERR
		errResult["errmsg"] = utils.RecodeText(utils.RECODE_DATAERR)
		ctx.JSON(200, errResult)
		return
	}
	if err != nil {
		errResult["errno"] = utils.RECODE_DATAERR
		errResult["errmsg"] = utils.RecodeText(utils.RECODE_DATAERR)
		ctx.JSON(200, errResult)
		return
	}
	grpcServer := utils.GetGrpcService()
	client := postLogin.NewPostLoginService("go.micro.srv.postLogin", grpcServer.Client())
	req := postLogin.Request{Mobile: logUser.Mobile, Password: logUser.Password}
	resp, err := client.Call(contextSys.TODO(), &req)
	if err != nil {
		fmt.Println(err)
		errResult["errno"] = utils.RECODE_LOGINERR
		errResult["errmsg"] = utils.RecodeText(utils.RECODE_LOGINERR)
		ctx.JSON(200, errResult)
		return
	}
	if resp.Errno==utils.RECODE_OK{
		s:=sessions.Default(ctx)
		s.Set("userName",resp.Name)
		err:=s.Save()
		if err!=nil{
			errResult["errno"]=utils.RECODE_SESSIONERR
			errResult["errmsg"]=utils.RecodeText(utils.RECODE_SESSIONERR)
			ctx.JSON(200,errResult)
		}
	}
	ctx.JSON(200, resp)

}
func GetUserInfo(ctx *gin.Context){
	s:=sessions.Default(ctx)
	userName:=s.Get("userName")
	grpcService:=utils.GetGrpcService()
	client:=getUserInfo.NewGetUserInfoService("go.micro.srv.getUserInfo",grpcService.Client())
	resp,err:=client.Call(contextSys.TODO(),&getUserInfo.Request{Name:userName.(string)})
	if err!=nil{
		fmt.Println(err)
		resp.Errno=utils.RECODE_DATAERR
		resp.Errmsg=utils.RecodeText(utils.RECODE_DATAERR)
		ctx.JSON(200,resp)
		return
	}
	resp.Errno=utils.RECODE_OK
	resp.Errmsg=utils.RecodeText(utils.RECODE_OK)
	ctx.JSON(200,resp)
}
