package utils

import (
	"fmt"
	"gojav/config"
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