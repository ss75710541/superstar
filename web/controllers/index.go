package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"log"
	"superstar/datasource"
	"superstar/models"
	"superstar/services"
)

type IndexController struct {
	Ctx iris.Context
	Service services.SuperstarService
}

// http://localhost:8080/
func (c *IndexController) Get() mvc.Result {
	//c.Service.GetAll()
	//return mvc.Response{
	//	Text:"ok\n",
	//}
	datalist := c.Service.GetAll()
	//var datalist []models.StarInfo
	//set the model and render the view template.
	return mvc.View{
		Name: "index.html",
		Data: iris.Map{
			"Title": "球星库",
			"Datalist": datalist,
		},
	}
}

func (c *IndexController) GetBy(id int) mvc.Result {
	if id < 1 {
		return mvc.Response{
			Path: "/",
		}
	}
	data := c.Service.Get(id)
	return mvc.View{
		Name: "info.html",
		Data: iris.Map{
			"Title": "球星库",
			"info": data,
		},
	}
}

func (c *IndexController) GetSearch() mvc.Result {
	country := c.Ctx.URLParam("country")
	datalist := c.Service.Search(country)
	return mvc.View{
		Name: "index.html",
		Data: iris.Map{
			"Title": "球星库",
			"Datalist": datalist,
		},
	}
}

// 集群多服务器的时候，才用得上这个接口
// 性能优化的时候才考虑，加上本机的SQL缓存
// http://localhost:8080/clearcache
func (c *IndexController) GetClearcache() mvc.Result {
	err := datasource.InstanceMaster().ClearCache(&models.StarInfo{})
	if err != nil {
		log.Fatalln(err)
	}

	// set the model and render the view template.
	return mvc.Response{
		Text: "xorm缓存清除成功",
	}
}