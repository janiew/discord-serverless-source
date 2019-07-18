package main


import (
	"context"
	"flag"
	"log"

	"github.com/PratikMahajan/Twitter-Knative-Serverless-App-Source/config"
	"github.com/knative/pkg/signals"
)

var (
	sink  string
	query string
)

func init() {
	flag.StringVar(&sink, "sink", "", "where to sink events to")
	query = config.MustGetEnvVar("QUERY", "")
}

func main() {

	flag.Parse()

	ctx := context.Background()

	// set up signals so we handle the first shutdown signal gracefully
	stopCh := signals.SetupSignalHandler()

	if query == "" {
		log.Fatal("Query parameter required")
	}

	log.Printf("Start (sink: %s, query: %s)", sink, query)

	search(ctx, query, sink, stopCh)

	<-stopCh
}