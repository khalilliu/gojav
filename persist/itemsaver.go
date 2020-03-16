package persist

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"log"
)

func ItemSaver() chan interface{} {
	out := make(chan interface{})

	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: Got %d item : %v\n", itemCount, item)
			itemCount++

			_, err := save(item)
			if err != nil {
				log.Printf("Item Saver :error saving item %v : %v ", item, err)
			}
		}
	}()

	return out
}

func save(item interface{})(string,  error) {
	client, err := elastic.NewClient(elastic.SetSniff(false))

	if err != nil {
		log.Println(err)
	}

	resp, err := client.Index().
		Index("datint_profile").
		Type("zhenai").
		BodyJson(item).
		Do(context.Background())

	if err != nil {
		log.Println(err)
		return "", err
	}

	fmt.Printf("%+v",resp)
	return resp.Id, nil
}
