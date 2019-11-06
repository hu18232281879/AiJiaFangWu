package model

import (
	"github.com/gomodule/redigo/redis"
	"microProject/getImagecd/utils"
)

func InitRedis() *redis.Pool{
	redisPool := &redis.Pool{
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp", utils.Redis_Address+":"+utils.Redis_Port)
		},
		MaxIdle:     10,
		MaxActive:   20,
		IdleTimeout: 20,
	}
	return redisPool
}
