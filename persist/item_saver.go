package persist

import (
	"bytes"
	"fmt"
	"gojav/config"
	"gojav/fetcher"
	"gojav/model"
	"gojav/utils"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func SaveItem(item model.Movie) {
	outputDir := config.Cfg.Output + "/" + item.Fanhao
	jsonFilePath := outputDir + "/" + item.Fanhao + ".json"
	if utils.IsExist(jsonFilePath) {
		fmt.Printf("文件: %s 已存在\n", jsonFilePath)
	}
	// save json
	itemSave := map[string]string{}
	mapMovie(item, &itemSave)
	utils.SaveFileToJson(itemSave, jsonFilePath)

	// save cover
	converFilePath := outputDir + "/" + item.Fanhao + ".jpg"
	if !utils.IsExist(converFilePath) {
		fmt.Printf("[%s][正在下载封面]: %s\n", item.Fanhao, item.Img)
		file, _ := os.Create(converFilePath)
		resp, err := http.Get(item.Img)
		defer resp.Body.Close()
		if err != nil {
			log.Fatal("save cover err: ", err)
		}
		body, _ := ioutil.ReadAll(resp.Body)
		io.Copy(file, bytes.NewReader(body))
	}

	for _, v := range item.Snapshot {
		wg.Add(1)
		go getSnapshot(item.Fanhao, v)
	}
	wg.Wait()
	return
}

func mapMovie(item model.Movie, itemSave *map[string]string) {
	(*itemSave)["番号"] = item.Fanhao
	(*itemSave)["影片地址"] = item.Link
	(*itemSave)["片名"] = item.Title
	(*itemSave)["发售日期"] = item.Date
	(*itemSave)["时长"] = item.Duration
	(*itemSave)["演员表"] = strings.Join(item.Actress, ",")
	(*itemSave)["类别"] = item.Series
	(*itemSave)["系列"] = strings.Join(item.Category, ",")
	if config.Cfg.Allmag {
		s := ""
		for i, v := range item.Magnets {
			if i != 0 {
				s += ","
			}
			s += fmt.Sprintf("%s:%s\n", v.SizeText, v.Link)
		}
		s = "[" + s + "]"
		(*itemSave)["下载磁链"] = strings.Join(item.Category, ",")
	} else {
		(*itemSave)["下载磁链"] = fmt.Sprintf("%s:%s", item.Magnets[0].SizeText, item.Magnets[0].Link)
	}
}

func getSnapshot(fanhao, url string) {
	strs := strings.Split(url, "/")
	name := strs[len(strs)-1]
	snapshotPath := fmt.Sprintf("%s/%s/%s%s", config.Cfg.Output, fanhao, fanhao, name)
	if !utils.IsExist(snapshotPath) {
		fmt.Printf("[%s][正在下载截图]: %s\n", fanhao, url)
		data, err := fetcher.FetchWithoutEncoding(url)
		if err != nil {
			log.Fatal("[getSnapshot] err : ", err)
			return
		}
		file, err := os.Create(snapshotPath)
		defer file.Close()
		if err != nil {
			panic(err)
		}
		file.Write(data)
	}
	defer wg.Done()
}
