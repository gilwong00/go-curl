package config

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type RequestConfig struct {
	Headers            http.Header
	UserAgent          string
	Data               string
	Method             string
	Insecure           bool
	Url                *url.URL
	ControlOutput      io.Writer
	ResponseBodyOutput io.Writer
}

func NewRequetConfig() *RequestConfig {
	return &RequestConfig{
		Headers:            map[string][]string{},
		ResponseBodyOutput: os.Stdout,
		ControlOutput:      os.Stdout,
	}
}

func (r *RequestConfig) AppendHeaders(headers []string) error {
	if len(headers) == 0 {
		return nil
	}
	for _, header := range headers {
		if key, value, found := strings.Cut(header, ":"); found {
			r.Headers.Add(strings.TrimSpace(key), strings.TrimSpace(value))
		} else {
			// TODO: clean this up
			// throw error
			return errors.New("invalid header")
		}
	}
	return nil
}
