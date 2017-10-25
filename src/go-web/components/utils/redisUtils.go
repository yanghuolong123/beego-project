package utils

import (
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
)

func Redis() (redis cache.Cache) {
	redis, err := cache.NewCache("redis", `{"conn":"127.0.0.1:6379", "dbNum":"0"}`)
	if err != nil {
		Logs().Critical(err.Error())
	}

	return
}
