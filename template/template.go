package template

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"text/template"

	"github.com/drone/funcmap"
)

// Render parses and executes a template, returning the results as string
// format. Trailing or leading spaces or new-lines are not getting truncated. It
// is able to read templates from remote paths, local files or directly from the
// string.
func Render(input string, payload interface{}) (string, error) {
	u, err := url.Parse(input)

	if err == nil {
		switch u.Scheme {
		case "http", "https":
			res, err := http.Get(input)

			if err != nil {
				return input, fmt.Errorf("failed to fetch: %w", err)
			}

			defer res.Body.Close()

			out, err := ioutil.ReadAll(res.Body)

			if err != nil {
				return input, fmt.Errorf("failed to read: %w", err)
			}

			input = string(out)
		case "file", "gotmpl":
			out, err := ioutil.ReadFile(u.Path)

			if err != nil {
				return input, fmt.Errorf("failed to read: %w", err)
			}

			input = string(out)
		}
	}

	tmpl, err := template.New("").Funcs(funcmap.Funcs).Parse(input)

	if err != nil {
		return input, fmt.Errorf("failed to parse: %w", err)
	}

	var (
		result = bytes.NewBufferString("")
	)

	if err := tmpl.Execute(result, payload); err != nil {
		return input, err
	}

	return result.String(), nil
}
