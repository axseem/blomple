package generate

import (
	"context"

	"github.com/axseem/blomple/database"
	"github.com/axseem/blomple/view"
)

func IndexPage(rootPath string, db *database.Queries) error {
	articles, err := db.ListArticles(context.Background())
	if err != nil {
		return err
	}

	return render(rootPath, "index", view.IndexPage(articles))
}
