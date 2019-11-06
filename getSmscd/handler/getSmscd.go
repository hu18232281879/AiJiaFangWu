package handler

import (
	"context"
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"math/rand"
	"microProject/getSmscd/model"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	getSmscd "microProject/getSmscd/proto/getSmscd"
	"microProject/getSmscd/utils"
	"strconv"
	"time"
)

type GetSmscd struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *GetSmscd) Call(ctx context.Context, req *getSmscd.Request, rsp *getSmscd.Response) error {
	phone := req.Phone
	uuid := req.Uuid
	text := req.Text
	redisPoll:=model.InitRedis()
	conn:=redisPoll.Get()
	defer conn.Close()
	code,err:=redis.String( conn.Do("get",uuid))
	if err!=nil{
		rsp.Errno=utils.RECODE_DATAERR
		rsp.Errmsg=utils.RecodeText(utils.RECODE_DATAERR)
		return err
	}
	if code!=text{
		rsp.Errno=utils.RECODE_DBERR
		rsp.Errmsg=utils.RecodeText(utils.RECODE_DBERR)
		return errors.New("验证码不正确")
	}
	rand.Seed(time.Now().UnixNano())
	vcode:=rand.Intn(899999)+100000
	radNum:=strconv.Itoa(vcode)
	client, err := dysmsapi.NewClientWithAccessKey("cn-hangzhou", "LTAI4FkgoSqad9KWK9sqcUfD", "ArTgvcOa2Rx0LPEEp4ewNJmCcLpxQY")

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.PhoneNumbers = phone
	request.SignName = "优尚之家"
	request.TemplateCode = "SMS_176530738"
	request.TemplateParam = `{"code":`+radNum+`}`

	response, err := client.SendSms(request)
	if err != nil {
		rsp.Errno=utils.RECODE_DATAERR
		rsp.Errmsg=utils.RecodeText(utils.RECODE_DATAERR)
		return err
	}
	if response.IsSuccess() {
		rsp.Errno=utils.RECODE_OK
		rsp.Errmsg=utils.RecodeText(utils.RECODE_OK)
		_,err=conn.Do("setex",phone+"_code",5*60,radNum)
		if err!=nil{
			rsp.Errno=utils.RECODE_DATAERR
			rsp.Errmsg=utils.RecodeText(utils.RECODE_DATAERR)
			return err
		}
	}else {
		rsp.Errno=utils.RECODE_SMSERR
		rsp.Errmsg=utils.RecodeText(utils.RECODE_SMSERR)
		return err
	}
	fmt.Printf("response is %#v\n", response)
	return nil
}
