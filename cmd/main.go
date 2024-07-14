package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func validateArgs(cmd *cobra.Command, args []string) error {
	if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
		return err
	}

	requestFlag, _ := cmd.Flags().GetString("request")

	switch requestFlag {
	case "GET", "POST", "PUT", "DELETE":
	default:
		return fmt.Errorf("invalid request method: %s", requestFlag)
	}

	return nil
}

var rootCmd = &cobra.Command{
	Use:   "gocurl <url>",
	Short: "A curl command line tool",
	Long:  `A curl command line tool that to make HTTP requests.`,
	Args:  validateArgs,
	Example: `
gocurl http://example.com/get
gocurl http://example.com -H 'Authorization: Bearer <token>'
gocurl http://example.com/post -X POST -d '{"key": "value"}' -H "Content-Type: application/json"
gocurl http://example.com/put -X PUT -d '{"key": "value"}' -H "Content-Type: application/json"`,
	Run: func(cmd *cobra.Command, args []string) {
		// verbose, _ := cmd.Flags().GetBool("verbose")
		// requestFlag, _ := cmd.Flags().GetString("request")
		// dataFlag, _ := cmd.Flags().GetString("data")
		// headers, _ := cmd.Flags().GetStringArray("header")
		fmt.Println(">>> heheh")

	},
}

func main() {
	rootCmd.Execute()
}
