package engine

import (
	"github.com/fatih/color"
	"gojav/fetcher"
	"log"
)

var (
	co = color.New(color.FgCyan)
)

type SimpleEngine struct {

}


func (e SimpleEngine)Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		/*
		c.Printf("Fetching %s\n", r.Url)
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetcher: error fetching url %s %v", r.Url, err)
			continue
		}
		parseResult := r.ParseFunc(body)
		 */

		parseResult, err := worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Printf(co.Sprintf("Got item %v",item))
		}
	}
}

func worker(r Request) (ParseResult, error) {
	co.Printf("Fetching %s\n", r.Url)

	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error fetching url %s %v", r.Url, err)
		return ParseResult{}, err
	}
	return r.ParseFunc(body), nil
}