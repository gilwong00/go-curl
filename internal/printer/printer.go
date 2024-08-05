package printer

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Printer struct {
	prefix string
	strings.Builder
}

func NewPrinter(prefix string) *Printer {
	return &Printer{
		prefix: ">",
	}
}

func (p *Printer) WriteResponse(
	res *http.Response,
	verbose bool,
	requestWriter io.Writer,
	bodyWriter io.Writer,
) error {
	if verbose {
		p.Printf("%v %v", res.Proto, res.Status)
		p.WriteHeaders(res.Header)
		p.Printf("")
		p.Println()
		if _, err := io.Copy(requestWriter, strings.NewReader(p.String())); err != nil {
			return err
		}
	}
	_, err := io.Copy(bodyWriter, res.Body)
	return err
}

func (p *Printer) WriteHeaders(headers http.Header) {
	for key, values := range headers {
		for _, value := range values {
			p.Printf("%v: %v", key, value)
		}
	}
}

func (w *Printer) Println() {
	w.WriteString("\n")
}

func (p *Printer) Printf(s string, a ...any) {
	p.WriteString(fmt.Sprintf("%v %v\n", p.prefix, fmt.Sprintf(s, a...)))
}
