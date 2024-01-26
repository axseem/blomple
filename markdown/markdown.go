package markdown

import (
	"bytes"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
)

func ConvertToHTML(s string) (string, error) {
	var html bytes.Buffer
	md := goldmark.New(
		goldmark.WithExtensions(
			extension.NewTable(
				extension.WithTableCellAlignMethod(
					extension.TableCellAlignAttribute,
				),
			),
			extension.GFM,
		),
	)

	err := md.Convert([]byte(s), &html)
	if err != nil {
		return "", err
	}

	return html.String(), nil
}
