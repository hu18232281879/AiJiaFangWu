package handler

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"microProject/postLogin/model"
	postLogin "microProject/postLogin/proto/postLogin"
	"microProject/postLogin/utils"
)

type PostLogin struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *PostLogin) Call(ctx context.Context, req *postLogin.Request, rsp *postLogin.Response) error {
	db, err := model.InitDb()
	defer db.Close()
	if err != nil {
		return err
	}
	m5:=md5.New()
	m5.Write([]byte(req.Password))
	hashPassword:=hex.EncodeToString(m5.Sum(nil))
	err,name:=model.QueryUser(db,req.Mobile,hashPassword)
	if err!=nil{
		rsp.Errno=utils.RECODE_LOGINERR
		rsp.Errmsg=utils.RecodeText(utils.RECODE_LOGINERR)
	}else {
		rsp.Errno=utils.RECODE_OK
		rsp.Errmsg=utils.RecodeText(utils.RECODE_OK)
		rsp.Name=name
	}
	return nil
}
