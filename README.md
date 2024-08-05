# gocurl

gocurl is a command line curl-like tool for making HTTP requests written in Golang.

# usage

`gocurl <url> [flags] `

with flags as follows:

- `-H` headers to be sent with the request
- `-u` the user agent to be used for requests (defaults to gocurl)
- `-m` specify an HTTP method (allowed: GET, PUT, POST, PATCH, DELETE, defaults to GET)
- `-k` allows insecure server connections over HTTPS (defaults to false).
- `-v` enable verbose logging. Displays all headers and the body (defaults to false).
