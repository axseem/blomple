package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "modernc.org/sqlite"

	"github.com/axseem/blomple/database"
	"github.com/axseem/blomple/handler"
	"github.com/axseem/blomple/render"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	rootPath := "public"
	err := os.Mkdir(rootPath, 0755)
	if os.IsExist(err) {
		info, err := os.Stat(rootPath)
		if err != nil {
			log.Fatal("failed to stat path: ", err)
		} else if !info.IsDir() {
			log.Fatal("path is not a directory: ", rootPath)
		}
	}

	sqlite, err := sql.Open("sqlite", "./sqlite.db")
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}
	db := database.New(sqlite)

	if err := render.IndexPage(rootPath, db); err != nil {
		log.Fatal(err)
	}
	if err := render.Articles(rootPath, db); err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/", handler.RootHandler(db))
	r.Get("/article/{id}", handler.ArticleHandler(db))
	http.ListenAndServe(":1323", r)
}
