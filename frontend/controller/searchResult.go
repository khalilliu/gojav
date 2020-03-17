package controller

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"gojav/engine"
	"gojav/frontend/model"
	"gojav/frontend/view"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type SearchResultHandler struct {
	view   view.SearchResultView
	client *elastic.Client
}

//localhost:9527/search?q=女&from=10
func (h SearchResultHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	q := strings.TrimSpace(r.FormValue("q"))
	from, err := strconv.Atoi(r.FormValue("from"))
	if err != nil {
		from = 0
	}

	fmt.Printf("q:%s, form:%d\n", q, from)
	result , err := h.getSearchResult(q, from)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.view.Render(w, result)

	fmt.Println("...", err)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (h SearchResultHandler) getSearchResult(q string, from int) (model.SearchResult, error) {
	var result model.SearchResult

	result.Query = q
	q = rewriteQueryString(q)

	resp, err := h.client.Search("datint_profile").
		Query(elastic.NewQueryStringQuery(q)).
		From(from).
		Do(context.Background())

	if err != nil {
		return result, err
	}
	fmt.Println("resp-->", resp)
	result.Hits = resp.TotalHits()
	result.Start = from

	itemsRaw := resp.Each(reflect.TypeOf(engine.Item{}))
	for _, v := range itemsRaw {
		item := v.(engine.Item)
		result.Items = append(result.Items, item)
	}
	result.PrevFrom = result.Start - len(result.Items)
	result.NextFrom = result.Start + len(result.Items)
	return result, nil
}

func CreateSearchResultHandler(template string) SearchResultHandler {
	client, err  := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	return SearchResultHandler{
		view: view.CreateSearchResultView(template),
		client:client,
	}
}


func rewriteQueryString(q string) string{
	re:=regexp.MustCompile(`([A-Z][a-z]*):`)
	s := re.ReplaceAllString(q,"Payload.$1:")
	fmt.Println("search for: -->",s)
	return s
}
