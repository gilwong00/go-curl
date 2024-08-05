package command

import (
	"net/http"

	"github.com/gilwong00/go-curl/internal/config"
	"github.com/gilwong00/go-curl/internal/httpclient"

	"github.com/spf13/cobra"
)

func CreateRootCommand() *cobra.Command {
	requestConfig := config.NewRequetConfig()
	headers := make([]string, 0, 255)
	var rootCmd = &cobra.Command{
		Use:   "gocurl <url>",
		Short: "A curl command line tool",
		Long:  `A curl command line tool used to make HTTP requests.`,
		Args:  requestConfig.ValidateArgs,
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
			return httpclient.ExecuteRequest(requestConfig)
		},
	}
	rootCmd.PersistentFlags().StringArrayVarP(&headers, "headers", "H", []string{}, `headers to be sent with the request, headers are separated by "," e.g. "Header1: value, Header2: Some other value"`)
	rootCmd.PersistentFlags().StringVarP(&requestConfig.UserAgent, "user-agent", "u", "gocurl", "the user agent to be used for requests")
	rootCmd.PersistentFlags().StringVarP(&requestConfig.Data, "data", "d", "", "data to be sent as the request body")
	rootCmd.PersistentFlags().StringVarP(&requestConfig.Method, "method", "m", http.MethodGet, "HTTP method to be used for the request")
	rootCmd.PersistentFlags().BoolVarP(&requestConfig.Insecure, "insecure", "k", false, "allows insecure server connections over HTTPS")
	rootCmd.PersistentFlags().BoolVarP(&requestConfig.Verbose, "verbose", "v", false, "can often be useful for debugging the request and generating documentation")
	return rootCmd
}
