package engine

import (
	"github.com/fatih/color"
	"gojav/fetcher"
	"log"
)

func Run(seeds ...Request) {
	c := color.New(color.FgCyan, color.Bold)
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		c.Printf("Fetching %s\n", r.Url)
		body, err := fetcher.Fetch(r.Url)
		if err != nil {
			log.Printf("Fetcher: error fetching url %s %v", r.Url, err)
			continue
		}
		parseResult := r.ParseFunc(body)
		requests = append(requests, parseResult.Requests...)
		for _, item := range parseResult.Items {
			log.Printf(c.Sprintf("Got item %v",item))
		}

	}
}