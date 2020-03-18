package fetcher

import (
	"bufio"
	"errors"
	"gojav/config"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"fmt"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Fetch(url string) ([]byte, error) {
	client := &http.Client{
		Timeout:  time.Duration(config.Cfg.Timeout  * 1e6),
	}

	request, _ := http.NewRequest(http.MethodGet,  url, nil)
	resp, err := client.Do(request)
	defer resp.Body.Close()
	if resp == nil {
		fmt.Println("resp error:", resp)
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Println("resp error:", resp)
		return nil, errors.New(fmt.Sprintf("error: status code: %d", resp.StatusCode))
	}

	bodyReader := bufio.NewReader(resp.Body)
	e := determinEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

func determinEncoding(r *bufio.Reader) encoding.Encoding {
	//bytes, e := bufio.NewReader(r).Peek(1024)
	bytes, e := bufio.NewReader(r).Peek(1024)
	if e != nil {
		//panic(e)
		log.Printf("Encoding error: %v", e)
		return unicode.UTF8
	}
	e2, _, _ := charset.DetermineEncoding(bytes, "")
	return e2
}