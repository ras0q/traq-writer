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

const (
	// DefaultHTTPOrigin is the default HTTP origin of traQ
	DefaultHTTPOrigin = "https://q.trap.jp"

	// webhookAPIPath is the webhook API path of traQ v3
	webhookAPIPath = "/api/v3/webhooks"
)

// TraqWebhookWriter implements io.Writer
type TraqWebhookWriter struct {
	id     string
	secret string
	origin string
}

// NewTraqWebhookWriter returns a new pointer of TraqWebhookWriter
func NewTraqWebhookWriter(id, secret, origin string) *TraqWebhookWriter {
	return &TraqWebhookWriter{id, secret, origin}
}

// Write posts a message to traQ via webhook
func (w *TraqWebhookWriter) Write(p []byte) (n int, err error) {
	url := fmt.Sprintf("%s%s/%s?embed=1", w.origin, webhookAPIPath, w.id)

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(p))
	if err != nil {
		return 0, fmt.Errorf("failed to create a new request: %w", err)
	}

	req.Header.Set("Content-Type", "text/plain; charset=utf-8")

	if w.isSecureMethod() {
		req.Header.Set("X-TRAQ-Signature", CalcHMACSHA1(w.secret, p))
	}

	httpClient := http.DefaultClient
	if _, err = httpClient.Do(req); err != nil {
		return 0, fmt.Errorf("failed to post a request: %w", err)
	}

	return len(p), nil
}

func (w *TraqWebhookWriter)isSecureMethod() bool {
	return len(w.secret) > 0
}

// CalcHMACSHA1 calculates an HMAC with SHA1
func CalcHMACSHA1(secret string, p []byte) string {
	mac := hmac.New(sha1.New, []byte(secret))
	mac.Write(p)

	return hex.EncodeToString(mac.Sum(nil))
}

// Interface guard
var _ io.Writer = (*TraqWebhookWriter)(nil)
