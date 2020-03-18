package fetcher

import (
	"gojav/config"
	"net/http"
	"net/url"
	"time"
)


var client *http.Client

func init() {
	var httpTransport *http.Transport
	if config.Cfg.Proxy != "" {
		proxyUrl, _ := url.Parse(config.Cfg.Proxy)
		httpTransport = &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
	} else  {
		httpTransport = &http.Transport{Proxy: http.ProxyFromEnvironment}
	}
	client = &http.Client{
		Transport: httpTransport,
		Timeout:  time.Duration(10 * time.Second),
	}
}

func Client() *http.Client {
	return client
}