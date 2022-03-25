package traqwriter

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"net/http"
)

// baseURL is a base API URL of traQ v3
const baseURL = "https://q.trap.jp/api/v3"

// TraqWebhookWriter implements TraqWebhookWriter
type TraqWebhookWriter struct {
	id     string
	secret string
}

// NewTraqWebhookWriter returns a new TraqWebhookWriter
func NewTraqWebhookWriter(id, secret string) *TraqWebhookWriter {
	return &TraqWebhookWriter{id, secret}
}

// Write posts a message to traQ via webhook
func (w *TraqWebhookWriter) Write(p []byte) (n int, err error) {
	url := fmt.Sprintf("%s/webhooks/%s?embed=1", baseURL, w.id)

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(p))
	if err != nil {
		return 0, fmt.Errorf("failed to create a new request: %w", err)
	}

	req.Header.Set("X-TRAQ-Signature", CalcHMACSHA1(w.secret, p))
	req.Header.Set("Content-Type", "text/plain; charset=utf-8")

	httpClient := http.DefaultClient
	if _, err = httpClient.Do(req); err != nil {
		return 0, fmt.Errorf("failed to post a request: %w", err)
	}

	return len(p), nil
}

// CalcHMACSHA1 calculates an HMAC with SHA1
func CalcHMACSHA1(secret string, p []byte) string {
	mac := hmac.New(sha1.New, []byte(secret))
	mac.Write(p)

	return hex.EncodeToString(mac.Sum(nil))
}

// Interface guard
var _ io.Writer = (*TraqWebhookWriter)(nil)
