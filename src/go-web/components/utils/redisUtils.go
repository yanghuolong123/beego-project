package utils

import (
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
)

var (
	Redis cache.Cache
)

func init() {
	Redis, _ = cache.NewCache("redis", `{"conn":"127.0.0.1:6379", "dbNum":"0"}`)
}
