package parser

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"gojav/config"
	"gojav/engine"
	"gojav/utils"
	"log"
	"strings"
)

func ParseMovieList(content [] byte) (result engine.ParseResult) {
	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(content))
	if err != nil {
		log.Fatal("error: [ParseMovieList]", err)
	}

	videoIds := []string{}

	nodes := doc.Find("a.movie-box")
	totalCountCurPage := nodes.Length()

	if config.Cfg.Limit != 0 && config.Cfg.Limit < totalCountCurPage {
		nodes = nodes.Slice(0, config.Cfg.Limit)
	}

	nodes.Each(func(i int, node *goquery.Selection) {
		if link, ok := node.Attr("href"); ok && link != "" {
			strs := strings.Split(link, "/")
			fanhao := strs[len(strs)-1]
			coverFilePath :=  config.Cfg.Output + "/" + fanhao + "/" + fanhao + ".jpg"
			magnetFilePath := config.Cfg.Output + "/" + fanhao + "/" + fanhao + ".json"

			if utils.IsExist(coverFilePath) && utils.IsExist(magnetFilePath) {
				log.Printf("[ %s ] Alreday fetched, SKIP!", strs[len(strs)-1]);
				return
			}
			videoIds = append(videoIds, strs[len(strs)-1])
			result.Requests = append(result.Requests, engine.Request{
				Url:link,
				Type: engine.HTML,
				ParseFunc: func(content []byte) engine.ParseResult {
					return ParseMovie(link, content)
				},
			})
			if num := config.Cfg.Limit; num != 0 {
				(&config.Cfg).Set("Limit", num-1)
			}
		}
	})

	fmt.Println("正在处理一下番号的影片...")
	fmt.Println(strings.Join(videoIds, ","))

	if config.Cfg.HasLimit && config.Cfg.Limit == 0  {
		engine.TargetHasFound = true
		return
	}

	//获取下一页
	nextPage, _ := doc.Find("#next").Eq(0).Attr("href")
	if len(nextPage) != 0 {
		nextLink := utils.GetUrl(nextPage)
		result.Requests = append(result.Requests, engine.Request{
			Url: nextLink,
			Type: engine.HTML,
			ParseFunc: ParseMovieList,
		})
	}
	return
}
