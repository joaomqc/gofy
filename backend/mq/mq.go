package mq

import (
	"encoding/json"
	"fmt"
	"gofy/config"
	"gofy/consumers"
	"gofy/infrastructure"
	"log"
	"time"

	rmq "github.com/adjust/rmq/v4"
)

func Init() {
	config := config.GetConfig()

	errChan := make(chan error)
	connection, err := rmq.OpenConnection("gofy", "tcp", config.GetString("mq.address"), 1, errChan)
	if err != nil {
		panic(err)
	}

	err = infrastructure.InitQueues(connection)
	if err != nil {
		panic(err)
	}

	err = initConsumers(connection)
	if err != nil {
		panic(err)
	}

	go handleErrors(errChan)
}

func handleErrors(errChan chan error) {
	for err := range errChan {
		switch e := err.(type) {
		case *rmq.HeartbeatError:
			fmt.Println("[MQ][HeartbeatError]", e)
		case *rmq.ConsumeError:
			fmt.Println("[MQ][ConsumeError]", e)
		case *rmq.DeliveryError:
			fmt.Println("[MQ][DeliveryError]", e)
		}
	}
}

func initConsumers(connection rmq.Connection) error {
	err := infrastructure.LoadArtistQueue.StartConsuming(10, time.Second)
	if err != nil {
		return err
	}
	artistHandler := consumers.LoadArtistConsumer{}
	infrastructure.LoadArtistQueue.AddConsumerFunc("load-artist-consumer", func(delivery rmq.Delivery) {
		handleMessage(delivery, artistHandler.HandleLoadArtistMessage)
	})

	err = infrastructure.LoadAlbumQueue.StartConsuming(10, time.Second)
	if err != nil {
		return err
	}
	albumHandler := consumers.LoadAlbumConsumer{}
	infrastructure.LoadAlbumQueue.AddConsumerFunc("load-album-consumer", func(delivery rmq.Delivery) {
		handleMessage(delivery, albumHandler.HandleLoadAlbumMessage)
	})

	err = infrastructure.DownloadTrackQueue.StartConsuming(10, time.Second)
	if err != nil {
		return err
	}
	trackHandler := consumers.DownloadTrackConsumer{}
	infrastructure.DownloadTrackQueue.AddConsumerFunc("download-track-consumer", func(delivery rmq.Delivery) {
		handleMessage(delivery, trackHandler.HandleDownloadTrackMessage)
	})

	return nil
}

func handleMessage[T any](delivery rmq.Delivery, handler func(T) error) {
	var message T
	if err := json.Unmarshal([]byte(delivery.Payload()), &message); err != nil {
		log.Print(err)
		if err := delivery.Reject(); err != nil {
			log.Print(err)
			// TODO handle reject error
		}
		return
	}
	fmt.Printf("handling message %T\n", message)
	if err := handler(message); err != nil {
		log.Print(err)
		if err := delivery.Reject(); err != nil {
			log.Print(err)
			// TODO handle reject error
		}
		return
	}

	if err := delivery.Ack(); err != nil {
		log.Print(err)
		// TODO handle ack error
	}
}
