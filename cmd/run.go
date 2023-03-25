package cmd

import (
	"encoding/gob"
	"fmt"
	"github.com/astaxie/beego"
	"mywork/models"
	"mywork/pkg/search"
	"mywork/routers"
	"strconv"
	"strings"
)

// Execute 初始化执行命令
func Execute() {

	routers.HomeApi()

	initSearch()

	tip()

}

func tip() {
	initSetSession()

	setTemplate()

	gob.Register(models.LiteOauthUser{})

	beego.Run("127.0.0.1:8089")
}

func setTemplate() {

	//创建新的模板函数
	err := beego.AddFuncMap("equrl", func(x, y string) bool {

		//去除获取url中的 /
		x1 := strings.Trim(x, "/")

		y1 := strings.Trim(y, "/")

		if strings.Compare(x1, y1) == 0 {

			return true
		} else {

			return false
		}

	})
	if err != nil {
		return
	}

	//拼接参数
	err = beego.AddFuncMap("urljoin", func(x string, y int) string {

		str := strconv.Itoa(y)
		return x + str
	})
	if err != nil {
		return
	}

	/**
	判读一个参数是否存在另一个集合里面
	*/
	err = beego.AddFuncMap("in_array", func(x uint, sli []string) bool {

		str := strconv.Itoa(int(x))

		for _, val := range sli {

			if val == str {

				return true
				break
			}
		}

		return false
	})
	if err != nil {
		return
	}

}

/**
初始化session
*/
func initSetSession() {

	beego.BConfig.WebConfig.Session.SessionOn = true

	//设置 cookies 的名字，Session 默认是保存在用户的浏览器 cookies 里面的，默认名是 beegosessionID，配置文件对应的参数名是
	beego.BConfig.WebConfig.Session.SessionName = "liteblog"

	beego.BConfig.WebConfig.Session.SessionProvider = "file"

	beego.BConfig.WebConfig.Session.SessionProviderConfig = "docs/session"
}

//初始化搜索引擎
func initSearch() {

	fmt.Println("初始化搜索引擎配置")
	host := beego.AppConfig.String("searchhost")
	port := beego.AppConfig.String("searchport")
	key := beego.AppConfig.String("searchkey")

	ser := NewMeiliSearch(host, port, key)

	search.SearchSDK = ser.OpenClient() //创建全局使用变量

	fmt.Println("init search success...")

}
