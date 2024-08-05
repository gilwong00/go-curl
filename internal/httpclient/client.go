package httpclient

import (
	"bytes"
	"crypto/tls"
	"io"
	"net/http"

	"github.com/gilwong00/go-curl/internal/config"
	"github.com/gilwong00/go-curl/internal/printer"
	"github.com/rs/zerolog/log"
)

func ExecuteRequest(c *config.RequestConfig) error {
	var reader io.Reader
	var tlsConfig *tls.Config
	// if we have a body append
	// should only append for non GET request
	if c.Data != "" {
		reader = bytes.NewBufferString(c.Data)
	}
	if c.Insecure {
		tlsConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}
	// create request
	request, err := http.NewRequest(c.Method, c.Url.String(), reader)
	if err != nil {
		return err
	}
	// append user agent if exists
	if c.UserAgent != "" {
		request.Header.Set("User-Agent", c.UserAgent)
	}
	// append headers
	for key, values := range c.Headers {
		for _, value := range values {
			request.Header.Add(key, value)
		}
	}
	// make http client
	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	res, err := client.Do(request)
	if err != nil {
		return err
	}
	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Warn().Err(err).Str("url", c.Url.String()).Msg("failed to close response body")
		}
	}()
	responseBuilder := printer.NewResponseBuilder(">")
	return responseBuilder.WriteResponse(res, c.Verbose, c.ControlOutput, c.ResponseBodyOutput)
}
