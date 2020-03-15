package controller

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"github.com/astaxie/beego/logs"
	"gojav/model"
	"strings"
	"time"
)

type CrawlMovieController struct {
	beego.Controller
}


func (c *CrawlMovieController) CrawlMovie() {
	c.Ctx.WriteString("starting...")
	var movieInfo model.MovieInfo
	//连接到redis
	model.ConnectRedis("127.0.0.1:6379")

	sUrl := "https://movie.douban.com/subject/25827935/" //七月与安生

	//添加到队列中
	model.PushQueue(sUrl)

	for {
		length := model.GetQueueLength()
		if length == 0 {
			break
		}

		sUrl = model.PopQueue()
		if model.IsVisit(sUrl) {
			continue
		}
		resp := httplib.Get(sUrl)
		resp.Header("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64; rv:50.0) Gecko/20100101 Firefox/50.0")
		resp.Header("Cookie", `bid=gFP9qSgGTfA; __utma=30149280.1124851270.1482153600.1483055851.1483064193.8; __utmz=30149280.1482971588.4.2.utmcsr=douban.com|utmccn=(referral)|utmcmd=referral|utmcct=/; ll="118221"; _pk_ref.100001.4cf6=%5B%22%22%2C%22%22%2C1483064193%2C%22https%3A%2F%2Fwww.douban.com%2F%22%5D; _pk_id.100001.4cf6=5afcf5e5496eab22.1482413017.7.1483066280.1483057909.; __utma=223695111.1636117731.1482413017.1483055857.1483064193.7; __utmz=223695111.1483055857.6.5.utmcsr=douban.com|utmccn=(referral)|utmcmd=referral|utmcct=/; _vwo_uuid_v2=BDC2DBEDF8958EC838F9D9394CC5D9A0|2cc6ef7952be8c2d5408cb7c8cce2684; ap=1; viewed="1006073"; gr_user_id=e5c932fc-2af6-4861-8a4f-5d696f34570b; __utmc=30149280; __utmc=223695111; _pk_ses.100001.4cf6=*; __utmb=30149280.0.10.1483064193; __utmb=223695111.0.10.1483064193`)

		sMovieHtml, err := resp.String()
		if err != nil {
			logs.Error(err)
		}

		movieInfo.Movie_name = model.GetMovieName(sMovieHtml)
		if movieInfo.Movie_name != "" {
			movieInfo.Movie_director = model.GetMovieDirector(sMovieHtml)
			movieInfo.Movie_main_character = model.GetMovieMainCharacters(sMovieHtml)
			movieInfo.Movie_type = model.GetMovieGenre(sMovieHtml)
			movieInfo.Movie_on_time = model.GetMovieOnTime(sMovieHtml) // 上映时间：2016-09-14(中国大陆)
			movieInfo.Movie_grade = model.GetMovieGrade(sMovieHtml)
			movieInfo.Movie_span = model.GetMovieRunningTime(sMovieHtml)

			movieInfo.Movie_on_time = movieInfo.Movie_on_time[0:strings.Index(model.GetMovieOnTime(sMovieHtml), "(")] //上映时间：2016-09-14

			movieInfo.Create_time = time.Now().Format("2006-1-2 15:04:05")

			id, _ := model.AddMovie(&movieInfo)
			fmt.Println("id -->", id)
		}

		model.AddToSet(sUrl)
		urls := model.GetMovieUrls(sMovieHtml)
		for _, url := range urls {
			if !model.IsVisit(url) {
				model.PushQueue(url)
				c.Ctx.WriteString("<br>" + url + "</br>")
			}
		}

		time.Sleep(time.Second)
	}
	c.Ctx.WriteString("end of crawl!")
}
