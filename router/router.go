package router

import (
	"github.com/astaxie/beego"
	"gojav/controller"
)

func init(){
	beego.Router("/", &controller.MainController{})
	beego.Router("/crawlmovie", &controller.CrawlMovieController{}, "*:CrawlMovie")
}