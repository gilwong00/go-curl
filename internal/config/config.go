package config

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/spf13/cobra"
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
	Verbose            bool
}

func NewRequetConfig() *RequestConfig {
	return &RequestConfig{
		Headers:            map[string][]string{},
		ResponseBodyOutput: os.Stdout,
		ControlOutput:      os.Stdout,
	}
}

func (r *RequestConfig) ValidateArgs(cmd *cobra.Command, args []string) error {
	// Check if url is provided
	if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
		return err
	}
	// Parse url, if url is invalid throw error
	u, err := url.Parse(args[0])
	if err != nil {
		return fmt.Errorf("the URL provided is invalid: %v, err: %w", args[0], err)
	}
	r.Url = u
	return nil
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
