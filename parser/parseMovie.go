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
	"path"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var (
	gid_re  = regexp.MustCompile(`gid\s+=\s+(\d+)`)
	uc_re   = regexp.MustCompile(`uc\s+=\s+(\d+);`)
	img_re  = regexp.MustCompile(`img\s+=\s+&#39;(http.+\.jpg)&#39;;`)
	size_re = regexp.MustCompile(`([0-9.]+)([GMTK]B)`)
)

func ParseMovie(link string, content []byte) engine.ParseResult {

	var result engine.ParseResult

	strs := strings.Split(link, "/")
	fanhao := strs[len(strs)-1]
	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(content))
	if err != nil {
		log.Fatal("error: [ParseMovie]", err)
	}

	movie := model.Movie{
		Fanhao: fanhao,
		Link:   link,
		Lang:   "zh",
	}

	// 创建目录
	itemOutDir := config.Cfg.Output + "/" + fanhao
	utils.EnsureNestDir(itemOutDir)
	mgnetFilePath := path.Join(itemOutDir, fanhao+".json")

	script, _ := doc.Find("body script").Eq(2).Html()

	// 获取 uc, gid, img
	parseScript(script, &movie)
	if p := itemOutDir + "/" + fanhao + ".jpg"; !utils.IsExist(p) {
		result.Requests = append(result.Requests, engine.Request{
			Url: movie.Img,
			Type: engine.IMG,
			ParseFunc: func(content []byte) engine.ParseResult {
				return ParseImg(content, p)
			},
		})
	}

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

	// 获取截图链接
	doc.Find("a.sample-box").Each(func(i int, selection *goquery.Selection) {
		href, _ := selection.Attr("href")
		movie.Snapshot = append(movie.Snapshot, href)
		// getsnapshot path
		strs := strings.Split(href, "/")
		name := strs[len(strs)-1]
		p := itemOutDir + "/" + name
		if !utils.IsExist(p) {
			result.Requests = append(result.Requests, engine.Request{
				Url:href,
				Type: engine.IMG,
				ParseFunc: func(content []byte) engine.ParseResult {
					return ParseImg(content, p)
				},
			})
		}
	})


	// 如果没有存储json, 就请求获取磁链
	if !utils.IsExist(mgnetFilePath) {
		url := fmt.Sprintf("%s/ajax/uncledatoolsbyajax.php?gid=%s&lang=%s&img=%s&uc=%s&floor=%d", config.BaseUrl, movie.Gid, movie.Lang, movie.Img, movie.Uc, utils.RandInt(1e3))
		body, err := fetcher.FetchWithoutEncoding(url)
		if err != nil {
			log.Fatal(err)
			utils.DeleteFile(itemOutDir)
			return result
		}

		doc2, err := goquery.NewDocumentFromReader(bytes.NewBuffer(body))
		if err != nil {
			log.Fatal("error: [ParseMovie]", err)
		}

		// 有种子资源, 进行解析
		if doc2.Find("a").Length() >= 1 {
			var magSizes model.Magnets

			doc2.Find("a").Each(func(i int, s *goquery.Selection) {
				if ok := size_re.MatchString(strings.TrimSpace(s.Text())); ok {
					mlink, _ := s.Attr("href")
					sizeText := strings.TrimSpace(s.Text())
					size := text2size(sizeText)
					magSizes = append(magSizes, model.Magnet{mlink, size, sizeText})
				}
			})

			//排序
			sort.Sort(model.SortBySize{magSizes})
			movie.Magnets = append(movie.Magnets, magSizes...)
		}
		// 没有存储该片, limit--
		result.Items = append(result.Items, engine.Item{movie})
	}
	return result
}

func parseScript(script string, movie *model.Movie) {
	gid := gid_re.FindStringSubmatch(script)[1]
	uc := uc_re.FindStringSubmatch(script)[1]
	img := img_re.FindStringSubmatch(script)[1]

	movie.Gid = gid
	movie.Uc = uc
	movie.Img = img
}

func text2size(text string) float64 {
	match := size_re.FindStringSubmatch(strings.TrimSpace(text))
	num, _ := strconv.ParseFloat(match[1], 64)
	var unit float64
	switch match[2] {
	case "GB":
		unit = 1000.00
		break
	case "MB":
		unit = 1.00
		break
	default:
		unit = 0
	}
	return unit * num
}
