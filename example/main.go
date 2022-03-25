package main

import (
	"fmt"
	"os"

	traqwriter "github.com/Ras96/traq-writer"
)

func main() {
	w := traqwriter.NewTraqWebhookWriter(
		os.Getenv("TRAQ_WEBHOOK_ID"),
		os.Getenv("TRAQ_WEBHOOK_SECRET"),
	)

	// post a message to traQ via webhook
	fmt.Fprint(w, "Hello, World!")
}
