package handler

import (
	"net/http"
	"strconv"

	"github.com/axseem/blomple/database"
	"github.com/axseem/blomple/markdown"
	"github.com/axseem/blomple/render"
	"github.com/axseem/blomple/view"
	"github.com/go-chi/chi/v5"
)

func ArticleHandler(db *database.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s := chi.URLParam(r, "id")
		id, err := strconv.Atoi(s)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Wrong ID"))
			return
		}

		article, err := db.GetArticle(r.Context(), int64(id))
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("No article found"))
			return
		}

		html, err := markdown.ConvertToHTML(article.Content)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Failed to parse article content"))
			return
		}

		content := render.HTMLToComonent(html)
		page := view.Article(article, content)
		page.Render(r.Context(), w)
	}
}
