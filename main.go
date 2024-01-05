package main

import (
	"log"
	"net/http"

	"github.com/axseem/blomple/database"
	"github.com/axseem/blomple/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	db, err := database.ConnectDB("./dev.db")
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}

	// err = db.CreateArticle(context.Background(), database.CreateArticleParams{
	// 	Title: "Hello World",
	// 	Body:  "This is the first article",
	// })
	// if err != nil {
	// 	log.Fatal(err)
	// }

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", handler.RootHandler(db))
	r.Get("/article/{id}", handler.ArticleHandler(db))
	http.ListenAndServe(":1323", r)
}
