package main

import (
	"context"
	"log"
	"strings"

	"github.com/janiew/discord-serverless-source/config"
	"github.com/bwmarrin/discordgo"
	
)

var (
	accessToken    = strings.TrimSpace(config.MustGetEnvVar("D_ACCESS_TOKEN", ""))
	webhookName = config.MustGetEnvVar("WEBHOOK_NAME", "")
)

var w map[string]*discordgo.Webhook

func search(ctx context.Context, query, sink string, stop <-chan struct{}) {

	// discord client config
	dg, err := discordgo.New("Bot " + accessToken)

	sinker, err := newSinkPoster(sink)
	if err != nil {
		log.Fatalf("Error getting sinker: %v", err)
	}

	w = make(map[string]*discordgo.Webhook)

	dg.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {

		if m.Author.ID == s.State.User.ID {
			return
		}
	
		if m.Content == query {
			var web = w[m.ChannelID]
			var found = web != nil
			if !found {
				var cwebs,err = dg.ChannelWebhooks(m.ChannelID)
				if err != nil {
					log.Printf("Error getting channel webhooks: %v", err)
					return
				}
				for _, v := range cwebs {
					if v.Name == webhookName {
						w[m.ChannelID] = v
						web = v
						found = true
					}
				}
				if !found {
					web,err = dg.WebhookCreate(m.ChannelID,webhookName,"")
					if err != nil {
						log.Printf("Error getting webhook: %v", err)
						return
					}

					w[m.ChannelID] = web
				}
			}

			sinker.post(ctx,s,m,web)
		}
	})

	err = dg.Open()
	if err != nil {
		log.Fatalf("Error opening connection to discord: %v\n", err)
	}

	log.Printf("Starting discord stream for: %s\n", query)
	dg.Close()

}


