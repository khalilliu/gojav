package parser

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"gojav/config"
	"gojav/engine"
	"gojav/utils"
	"log"
	"path"
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
	if config.Cfg.Limit != 0 && config.Cfg.Limit > totalCountCurPage {
		nodes = nodes.Slice(0, config.Cfg.Limit)
	}

	nodes.Each(func(i int, node *goquery.Selection) {
		if link, ok := node.Attr("href"); ok && link != "" {
			strs := strings.Split(link, "/")
			videoIds = append(videoIds, strs[len(strs)-1])

			coverFilePath := path.Join(config.Cfg.Output, strs[len(strs)-1] + ".jpg")
			magnetFilePath := path.Join(config.Cfg.Output, strs[len(strs)-1] + ".jpg")

			if utils.IsExist(coverFilePath) && utils.IsExist(magnetFilePath) {
				log.Printf("[ %s ] Alreday fetched, SKIP!", strs[len(strs)-1]);
				return
			}

			result.Requests = append(result.Requests, engine.Request{
				Url:link,
				ParseFunc: func(content []byte) engine.ParseResult {
					return ParseMovie(link, content)
				},
			})
		}
	})
	fmt.Println("正在处理一下番号的影片...")
	fmt.Println(strings.Join(videoIds, ","))
	return
}
