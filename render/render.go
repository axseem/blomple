package render

import (
	"context"
	"io"
	"os"
	"path"

	"github.com/a-h/templ"
	"github.com/microcosm-cc/bluemonday"
)

func HTMLToComonent(html string) templ.Component {
	sanitizedHTML := bluemonday.UGCPolicy().Sanitize(html)
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) error {
		_, err := w.Write([]byte(sanitizedHTML))
		return err
	})
}

func render(pageDir string, name string, component templ.Component) error {
	pagePath := path.Join(pageDir, name+".html")
	f, err := os.Create(pagePath)
	if err != nil {
		return err
	}

	return component.Render(context.Background(), f)
}
