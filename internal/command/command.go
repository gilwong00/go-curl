package command

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/gilwong00/go-curl/internal/config"
	"github.com/gilwong00/go-curl/internal/request"

	"github.com/spf13/cobra"
)

func CreateRootCommand() *cobra.Command {
	requestConfig := config.NewRequetConfig()
	headers := make([]string, 0, 255)
	var rootCmd = &cobra.Command{
		Use:   "gocurl <url>",
		Short: "A curl command line tool",
		Long:  `A curl command line tool that to make HTTP requests.`,
		Args:  validateArgs,
		Example: `
	gocurl http://example.com/get
	gocurl http://example.com -H 'Authorization: Bearer <token>'
	gocurl http://example.com/post -m POST -d '{"key": "value"}' -H "Content-Type: application/json"
	gocurl http://example.com/put -m PUT -d '{"key": "value"}' -H "Content-Type: application/json"`,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			requestConfig.AppendHeaders(headers)
			return nil
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return request.ExecuteRequest(requestConfig)
		},
	}
	rootCmd.PersistentFlags().StringArrayVarP(&headers, "headers", "H", []string{}, `headers to be sent with the request, headers are separated by "," e.g. "Header1: value, Header2: Some other value"`)
	rootCmd.PersistentFlags().StringVarP(&requestConfig.UserAgent, "user-agent", "u", "gocurl", "the user agent to be used for requests")
	rootCmd.PersistentFlags().StringVarP(&requestConfig.Data, "data", "d", "", "data to be sent as the request body")
	rootCmd.PersistentFlags().StringVarP(&requestConfig.Method, "method", "m", http.MethodGet, "HTTP method to be used for the request")
	rootCmd.PersistentFlags().BoolVarP(&requestConfig.Insecure, "insecure", "k", false, "allows insecure server connections over HTTPS")

	return rootCmd
}

func validateArgs(cmd *cobra.Command, args []string) error {
	// Check if url is provided
	if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
		return err
	}
	// Parse url, if url is invalid throw error
	u, err := url.Parse(args[0])
	if err != nil {
		return fmt.Errorf("the URL provided is invalid: %v, err: %w", args[0], err)
	}
	// TODO: append url to requestConfig
	fmt.Println(">>>> u", u)
	return nil
}
