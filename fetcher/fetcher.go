package fetcher

import (
	"bufio"
	"errors"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
)

func Fetch(rUrl string) ([]byte, error) {
	request, _ := http.NewRequest(http.MethodGet,  rUrl, nil)
	request.Header.Set("Referer", "https://www.javbus.com")

	resp, err := HttpClient.Do(request)
	if err != nil {
		fmt.Println("resp error:", err)
		return nil, err
	}
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

func FetchWithoutEncoding(rUrl string) ([]byte, error) {
	request, _ := http.NewRequest(http.MethodGet,  rUrl, nil)
	request.Header.Set("Referer", "https://www.javbus.com")

	resp, err := HttpClient.Do(request)
	if err != nil {
		fmt.Println("resp error:", err)
		return nil, err
	}
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
	return ioutil.ReadAll(bodyReader)
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