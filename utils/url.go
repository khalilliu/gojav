package utils

import (
	"fmt"
	"gojav/config"
	"os"
)

func GetUrl(next string) string {
	url := config.BaseUrl
	if config.Cfg.Search != "" {
		url = fmt.Sprintf("%s%s/%s", url, config.SearchRoute, config.Cfg.Search)
	} else if config.Cfg.Base != "" {
		url = config.Cfg.Base
	}

	if len(next) != 0 {
		url += next
	}
	return url
}

func UserHome() string {
	dir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return dir
}
