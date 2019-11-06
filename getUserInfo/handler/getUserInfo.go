package handler

import (
	"context"
	"microProject/getUserInfo/model"
	"microProject/getUserInfo/utils"

	getUserInfo "microProject/getUserInfo/proto/getUserInfo"
)

type GetUserInfo struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *GetUserInfo) Call(ctx context.Context, req *getUserInfo.Request, rsp *getUserInfo.Response) error {
	db, err := model.InitDb()
	if err != nil {
		rsp.Errno = utils.RECODE_DBERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
		return nil
	}
	user, err := model.ObtainUserInfo(db, req.Name)
	if err != nil {
		rsp.Errno = utils.RECODE_SESSIONERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_SESSIONERR)
		return nil
	}
	var userInfo getUserInfo.UserInfo
	userInfo.UserId = int32(user.ID)
	userInfo.Name = user.Name
	userInfo.Mobile = user.Mobile
	userInfo.RealName = user.Real_name
	userInfo.IdCard=user.Id_card
	userInfo.AvatarUrl = user.Avatar_url
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(utils.RECODE_OK)
	rsp.Data = &userInfo
	return nil
}
