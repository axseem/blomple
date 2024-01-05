package handler

import (
	"net/http"

	"github.com/axseem/blomple/database"
	"github.com/axseem/blomple/view"
)

func RootHandler(db *database.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		articles, err := db.ListArticles(r.Context())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("No articles found"))
			return
		}

		page := view.Root(articles)
		page.Render(r.Context(), w)
	}
}
