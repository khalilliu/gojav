package persist

import (
	"context"
	"github.com/olivere/elastic"
	"github.com/pkg/errors"
	"gojav/engine"
	"log"
)

func ItemSaver() chan engine.Item {
	out := make(chan engine.Item)

	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: Got %d item : %v\n", itemCount, item)
			itemCount++

			err := save(item)
			if err != nil {
				log.Printf("Item Saver :error saving item %v : %v ", item, err)
			}
		}
	}()

	return out
}

func save(item engine.Item) error {
	client, err := elastic.NewClient(elastic.SetSniff(false))

	if err != nil {
		log.Println(err)
	}

	if item.Type == "" {
		return errors.New("must supply Type...")
	}

	indexService := client.Index().
		Index("datint_profile").
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
