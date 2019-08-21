package main

import (
	"context"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
	ce "github.com/knative/pkg/cloudevents"
)

func newSinkPoster(sink string) (sinker *sinkPoster, err error) {

	// cloud events
	ceClient := ce.NewClient(sink, ce.Builder{
	})

	s := &sinkPoster{
		client: ceClient,
	}



	return s, nil

}

type sinkPoster struct {
	client *ce.Client
}

type MessageEvent struct {
	Message *discordgo.MessageCreate `json:"message"`
	Webhook *discordgo.Webhook `json:"webhook"`
}

func (s *sinkPoster) post(ctx context.Context, sesh *discordgo.Session, m *discordgo.MessageCreate, webhook *discordgo.Webhook) error {

	log.Printf("Posting message: %s\n", m.ID)
	eventTime, err := m.Timestamp.Parse()
	if err != nil {
		log.Printf("Error while parsing created at: %v", err)
		eventTime = time.Now()
	}

	data := MessageEvent{
		Message : m,
		Webhook : webhook,
	}

	return s.client.Send(data, ce.V01EventContext{
		EventID:   m.ID,
		EventTime: eventTime,
	})

}