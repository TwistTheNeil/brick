package utils

import (
	"bytes"
	"fmt"
	"text/tabwriter"
)

// Tabulate some data to a string
func Tabulate(data string) (string, error) {
	var buf bytes.Buffer
	w := new(tabwriter.Writer)

	// Format in tab-separated columns with a tab stop of 8.
	w.Init(&buf, 0, 8, 0, '\t', 0)
	fmt.Fprintln(w, data)
	if err := w.Flush(); err != nil {
		return "", nil
	}

	return buf.String(), nil
}
