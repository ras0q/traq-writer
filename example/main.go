package main

import (
	"fmt"
	"os"

	traqwriter "github.com/ras0q/traq-writer"
)

func main() {
	w := traqwriter.NewTraqWebhookWriter(
		os.Getenv("TRAQ_WEBHOOK_ID"),
		os.Getenv("TRAQ_WEBHOOK_SECRET"),
		traqwriter.DefaultHTTPOrigin,
	)

	// post a message to traQ via webhook
	fmt.Fprint(w, "Hello, World!")
}
