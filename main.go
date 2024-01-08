package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "modernc.org/sqlite"

	"github.com/axseem/blomple/database"
	"github.com/axseem/blomple/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	sqlite, err := sql.Open("sqlite", "./sqlite.db")
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}
	db := database.New(sqlite)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", handler.RootHandler(db))
	r.Get("/article/{id}", handler.ArticleHandler(db))
	http.ListenAndServe(":1323", r)
}
