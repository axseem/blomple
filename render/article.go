package render

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"path"

	"github.com/axseem/blomple/database"
	"github.com/axseem/blomple/view"
	"github.com/yuin/goldmark"
)

func Articles(rootPath string, db *database.Queries) error {
	articles, err := db.ListArticles(context.Background())
	if err != nil {
		return err
	}

	articlesPath := path.Join(rootPath, "article")
	err = os.Mkdir(articlesPath, 0755)
	if os.IsExist(err) {
		info, err := os.Stat(articlesPath)
		if err != nil {
			log.Fatal("failed to stat path: ", err)
		} else if !info.IsDir() {
			log.Fatal("path is not a directory: ", articlesPath)
		}
	}

	for _, article := range articles {
		var html bytes.Buffer
		if err := goldmark.Convert([]byte(article.Content), &html); err != nil {
			return err
		}

		content := HTMLToComonent(html.String())
		if err := render(articlesPath, fmt.Sprint(article.ID), view.Article(article, content)); err != nil {
			return err
		}
	}

	return nil
}
