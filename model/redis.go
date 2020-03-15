package model

import (
	"github.com/astaxie/goredis"
	"log"
)

const (
	URL_QUEUE = "url_queue"
	URL_VISIT_SET = "url_visit_set"
)

var (
	client goredis.Client
)

func ConnectRedis(addr string) {
	client.Addr = addr
}

func PushQueue(url string) {
	client.Lpush(URL_QUEUE, []byte(url))
}

func PopQueue() string {
	res,err := client.Rpop(URL_QUEUE)
	if err != nil {
		log.Fatal(err)
	}
	return string(res)
}

func GetQueueLength() int {
	length, err := client.Llen(URL_QUEUE)
	if err != nil {
		return 0
	}
	return length
}

func AddToSet(url string) {
	client.Sadd(URL_VISIT_SET, []byte(url))
}

func IsVisit(url string) bool {
	visited, err := client.Sismember(URL_VISIT_SET, []byte(url))
	if err != nil {
		return false
	}
	return visited
}