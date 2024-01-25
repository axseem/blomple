package render

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/axseem/blomple/database"
	"github.com/axseem/blomple/view"
	"github.com/yuin/goldmark"
)

func Articles(rootPath string, db *database.Queries) error {
	articles, err := db.ListArticles(context.Background())
	if err != nil {
		return err
	}

	err = os.Mkdir(rootPath, 0755)
	if os.IsExist(err) {
		info, err := os.Stat(rootPath)
		if err != nil {
			log.Fatal("failed to stat path: ", err)
		} else if !info.IsDir() {
			log.Fatal("path is not a directory: ", rootPath)
		}
	}

	for _, article := range articles {
		var html bytes.Buffer
		if err := goldmark.Convert([]byte(article.Content), &html); err != nil {
			return err
		}

		content := HTMLToComonent(html.String())
		if err := render(rootPath, fmt.Sprint(article.ID), view.Article(article, content)); err != nil {
			return err
		}
	}

	return nil
}
