package handler

import (
	"context"
	"encoding/json"
	"github.com/afocus/captcha"
	"image/color"
	"microProject/getImagecd/model"
	getImagecd "microProject/getImagecd/proto/getImagecd"
	"microProject/getImagecd/utils"
)

type GetImagecd struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *GetImagecd) Call(ctx context.Context, req *getImagecd.Request, rsp *getImagecd.Response) error {
	cap := captcha.New()
	if err := cap.SetFont("./handler/comic.ttf"); err != nil {
		panic(err.Error())
	}
	cap.SetSize(128, 64)
	//设置图片大小
	cap.SetDisturbance(captcha.NORMAL)
	cap.SetFrontColor(color.RGBA{255, 255, 255, 255})
	cap.SetBkgColor(color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255}, color.RGBA{0, 153, 0, 255})
	img, str := cap.Create(6, captcha.NUM)
	imgBuffer, err := json.Marshal(img)
	if err != nil {
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DATAERR)
		rsp.Img = nil
		return nil
	}
	redisPoll := model.InitRedis()
	conn := redisPoll.Get()
	defer conn.Close()
	_, err = conn.Do("setex", req.Uuid, 60*5, str)
	if err != nil {
		rsp.Errno = utils.RECODE_DATAERR
		rsp.Errmsg = utils.RecodeText(utils.RECODE_DATAERR)
		rsp.Img = nil
		return nil
	}
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(utils.RecodeText(utils.RECODE_OK))
	rsp.Img = imgBuffer
	return nil
}
