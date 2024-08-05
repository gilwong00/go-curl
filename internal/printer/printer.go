package printer

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type ResponseBuilder struct {
	prefix string
	strings.Builder
}

func NewResponseBuilder(prefix string) *ResponseBuilder {
	return &ResponseBuilder{
		prefix: ">",
	}
}

func (r *ResponseBuilder) WriteResponse(
	res *http.Response,
	verbose bool,
	requestWriter io.Writer,
	bodyWriter io.Writer,
) error {
	if verbose {
		r.Printf("%v %v", res.Proto, res.Status)
		r.WriteHeaders(res.Header)
		r.Printf("")
		r.Println()
		if _, err := io.Copy(requestWriter, strings.NewReader(r.String())); err != nil {
			return err
		}
	}
	_, err := io.Copy(bodyWriter, res.Body)
	return err
}

func (w *ResponseBuilder) WriteHeaders(headers http.Header) {
	for key, values := range headers {
		for _, value := range values {
			w.Printf("%v: %v", key, value)
		}
	}
}

func (w *ResponseBuilder) Println() {
	w.WriteString("\n")
}

func (w *ResponseBuilder) Printf(s string, a ...any) {
	w.WriteString(fmt.Sprintf("%v %v\n", w.prefix, fmt.Sprintf(s, a...)))
}
