package engine

import (
	"github.com/fatih/color"
	"gojav/fetcher"
	"log"
)

func worker(r Request) (ParseResult, error) {
	var (
		body []byte
		err  error
	)
	switch r.Type {
	case IMG:
		color.Yellow("[正在下载图片]: %s\n", r.Url)
		body, err = fetcher.FetchWithoutEncoding(r.Url)
		break
	case HTML:
		body, err = fetcher.Fetch(r.Url)
		break
	}
	if err != nil {
		log.Printf("Fetcher: error fetching url %s %v", r.Url, err)
		return ParseResult{}, err
	}
	return r.ParseFunc(body), nil
}
