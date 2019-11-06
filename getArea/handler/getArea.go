package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"getArea/model"
	getArea "getArea/proto/getArea"
	"getArea/utils"
	"github.com/gomodule/redigo/redis"
)

type GetArea struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *GetArea) Call(ctx context.Context, req *getArea.Request, rsp *getArea.Response) error {
	redisPoll := model.InitRedis()
	redisConn := redisPoll.Get()
	defer redisConn.Close()
	areaBuffer, _ := redis.Bytes(redisConn.Do("get", "Sf_area"))
	var areas []model.Area
	if len(areaBuffer) == 0 {
		db, err := model.InitDb()
		areastemp, err := model.GetAllArea(db)
		if err != nil {
			rsp.Errno = utils.RECODE_DATAERR
			rsp.Errmsg = utils.RecodeText(utils.RECODE_DATAERR)
			return err
		}
		areas=areastemp
		tempBuffer, err := json.Marshal(areastemp)
		_, err = redisConn.Do("set", "Sf_area", tempBuffer)
		if err != nil {
			fmt.Println(err)
			return err
		}
	} else {
		err := json.Unmarshal(areaBuffer, &areas)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	//给返回数据赋值
	rsp.Errno = utils.RECODE_OK
	rsp.Errmsg = utils.RecodeText(utils.RECODE_OK)
	for _, v := range areas {
		temp := new(getArea.Area)
		temp.Aid = int32(v.Id)
		temp.Aname = v.Name
		rsp.Data = append(rsp.Data, temp)
	}
	return nil
}
