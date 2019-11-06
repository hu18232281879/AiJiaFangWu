package handler

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/gomodule/redigo/redis"
	"microProject/postRet/model"
	"microProject/postRet/utils"

	postRet "microProject/postRet/proto/postRet"
)

type PostRet struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *PostRet) Call(ctx context.Context, req *postRet.Request, rsp *postRet.Response) error {
	redisConn:=model.InitRedis().Get()
	defer redisConn.Close()
	smsCode,err:=redis.String( redisConn.Do("get",req.Mobile+"_code"))
	if err!=nil{
		rsp.Errno=utils.RECODE_DATAERR
		rsp.Errmsg=utils.RecodeText(utils.RECODE_DATAERR)
		return err
	}
	if smsCode!=req.SmsCode{
		rsp.Errno=utils.RECODE_DATAERR
		rsp.Errmsg=utils.RecodeText(utils.RECODE_DATAERR)
		return errors.New("短信验证码错误")
	}
	var user model.User
	user.Name=req.Mobile
	user.Mobile=req.Mobile
	m5:=md5.New()
	m5.Write([]byte(req.Password))
	hashPwd:=hex.EncodeToString(m5.Sum(nil))
	user.Password_hash=hashPwd
	db,err:=model.InitDb()
	defer db.Close()
	if err!=nil{
		rsp.Errno=utils.RECODE_DATAERR
		rsp.Errmsg=utils.RecodeText(utils.RECODE_DATAERR)
		return err
	}
	err=model.InsertUser(db,user)
	if err!=nil{
		rsp.Errno=utils.RECODE_DATAERR
		rsp.Errmsg=utils.RecodeText(utils.RECODE_DATAERR)
		return err
	}
	rsp.Errno=utils.RECODE_OK
	rsp.Errmsg=utils.RecodeText(utils.RECODE_OK)


	//添加session


	return nil
}

