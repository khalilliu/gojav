package persist

import (
	"context"
	"github.com/olivere/elastic"
	"github.com/pkg/errors"
	"gojav/engine"
	"log"
)

func ItemSaver(index string) (chan engine.Item, error) {
	client, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		log.Println(err)
		return nil, err
	}

	out := make(chan engine.Item)

	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: Got %d item : %v\n", itemCount, item)
			itemCount++

			err := save(client, item, index)
			if err != nil {
				log.Printf("Item Saver :error saving item %v : %v ", item, err)
			}
		}
	}()

	return out, nil
}

func save(client *elastic.Client, item engine.Item, index string) error {

	if item.Type == "" {
		return errors.New("must supply Type...")
	}

	indexService := client.Index().
		Index(index).
		Type(item.Type).
		BodyJson(item)

	if item.Id != "" {
		indexService.Id(item.Id)
	}

	_, err = indexService.Do(context.Background())

	if err != nil {
		return err
	}

	return nil
}
