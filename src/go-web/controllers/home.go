package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"go-web/components/ext"
	"go-web/components/utils"
	"go-web/models"
	"log"
	"os"
	"time"
)

var (
	globalNum int
)

func init() {
	globalNum = 0
}

type HomeController struct {
	//beego.Controller
	ext.BaseController
}

func (this *HomeController) Get() {
	user := this.GetSession("user")
	this.Data["user"] = user

	//utils.Logs().Info("1111weeee")
	utils.Log.Info("1111weeee")

	this.Layout = "layout/main.tpl"
	this.TplName = "home/index.tpl"
}

func (c *HomeController) TestGet() {
	//m := make(map[string]string)
	//m["username"] = "杨活龙"
	//m["email"] = "yhl27ml@163.com"
	//c.Data["json"] = m

	//c.ServeJSON()

	dbhost := beego.AppConfig.String("mysqlhost")
	//utils.Log.Info("==== mysqlhost:" + dbhost)
	//utils.Log.Info(dbhost)
	_ = dbhost

	m, err := models.GetUser("1")
	if err != nil {
		c.SendRes(-1, "failed", err)
	}
	_ = m
	o := orm.NewOrm()
	var user models.User
	err = o.Raw("select * from tbl_user where username=? and email=?", "yhl27ml@126.com", "yhl27ml@126.com").QueryRow(&user)
	if err != nil {
		c.SendRes(-1, err.Error(), nil)
	}

	u, err := models.GetById(4)
	if err != nil {
		c.SendRes(-1, err.Error(), nil)
	}
	b, _ := models.Login("yhl27ml@163.com", "654321")
	log.Println(b)
	redis := utils.Redis
	err = redis.Put("test1", "1111111", 60*time.Second)
	err = redis.Put("test2", "2222222222222", 60*time.Second)
	log.Println("err:", err)
	log.Println("test1:", string(redis.Get("test1").([]uint8)))

	y, month, d := utils.Date()
	log.Println("y/m/d:", y, month, d)
	dir := fmt.Sprintf("%d/%d/%d", y, month, d)
	dir = "static/uploads/" + dir
	os.MkdirAll(dir, os.ModePerm)
	log.Println(dir)
	log.Println(int(time.November))
	log.Println(time.Now().Format(utils.DatetimeFormat))
	log.Println("Exist:", utils.PathExist("static/uploads/tmp"))
	globalNum++
	log.Println("======= globalNum:", globalNum)
	log.Println("gg=", string(redis.Get("gg").([]byte)))

	c.SendRes(0, "sucess", u)
	//c.SendRes(0, "sucess", models.GetAllUser())
}

func (c *HomeController) TestPost() {
	email := c.GetString("email")
	passwd := c.GetString("password")
	c.Ctx.WriteString(email + ":" + passwd)
}
