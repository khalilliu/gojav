package parser

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"gojav/config"
	"gojav/engine"
	"gojav/fetcher"
	"gojav/model"
	"gojav/utils"
	"log"
	"os"
	"path"
	"regexp"
	"strings"
)

func ParseMovie(link string, content []byte) engine.ParseResult {
	strs := strings.Split(link, "/")
	fanhao := strs[len(strs)-1]
	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(content))
	if err != nil {
		log.Fatal("error: [ParseMovie]", err)
	}

	movie := model.Movie{
		Fanhao: fanhao,
		Link: link,
		Lang: "zh",
	}

	script, _ := doc.Find("body script").Eq(2).Html()

	parseScript(script, &movie)

	title := doc.Find("div h3").Eq(0).Text()
	movie.Title = title

	doc.Find("div.info > p").Each(func(i int, selection *goquery.Selection) {
		text := selection.Text()
		if strings.Contains(text, "發行日期:") {
			movie.Date = strings.Replace(text, "發行日期: ", "", 1)
		} else if strings.Contains(text, "系列: ") {
			movie.Series = strings.Replace(text, "系列:", "", 1)
		} else if strings.Contains(text, "長度:") {
			movie.Duration = strings.Replace(text, "長度: ", "", 1)
		}

		selection.Find(".genre > a").Each(func(i int, selection *goquery.Selection) {
			t := selection.Text()
			movie.Category = append(movie.Category, t)
		})
	})
	//提取演员
	doc.Find("#avatar-waterfall > a span").Each(func(i int, selection *goquery.Selection) {
		t := selection.Text()
		movie.Actress = append(movie.Actress, t)
	})

	itemOutDir := config.Cfg.Output + "/" + fanhao
	// 创建目录
	utils.EnsureNestDir(itemOutDir)
	mgnetFilePath := path.Join(itemOutDir, fanhao + ".json")
	if !utils.IsExist(mgnetFilePath) {
		url := fmt.Sprintf("%s/ajax/uncledatoolsbyajax.php?gid=%s&lang=%s&img=%s&uc=%s&floor=%d", config.BaseUrl, movie.Gid, movie.Lang, movie.Img, movie.Uc, utils.RandInt(1e3) )
		body, err := fetcher.Fetch(url)
		if err != nil {
			log.Fatal(err)
			os.Remove(itemOutDir)
			return engine.ParseResult{}
		}



	}



	if num := config.Cfg.Limit; num != 0 {
		(&config.Cfg).Set("Limit", num-1)
	}

	return engine.ParseResult{}
}

func parseScript(script string, movie *model.Movie) {
	var (
		gid_re = regexp.MustCompile(`gid\s+=\s+(\d+)`)
		uc_re  = regexp.MustCompile(`uc\s+=\s+(\d+);`)
		img_re = regexp.MustCompile(`img\s+=\s+&#39;(http.+\.jpg)&#39;;`)
	)
	gid := gid_re.FindStringSubmatch(script)[1]
	uc := uc_re.FindStringSubmatch(script)[1]
	img := img_re.FindStringSubmatch(script)[1]

	movie.Gid = gid
	movie.Uc = uc
	movie.Img = img
}
