package fetcher

import (
	"gojav/config"
	"net/http"
	"net/url"
)


var HttpClient *http.Client

func init() {
	var httpTransport *http.Transport
	if config.Cfg.Proxy != "" {
		proxyUrl, _ := url.Parse(config.Cfg.Proxy)
		httpTransport = &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
	} else  {
		httpTransport = &http.Transport{Proxy: http.ProxyFromEnvironment}
	}
	HttpClient = &http.Client{
		Transport: httpTransport,
	}
}

func Client() *http.Client {
	return HttpClient
}