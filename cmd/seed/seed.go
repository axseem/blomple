package main

import (
	"context"
	"database/sql"
	"log"

	_ "modernc.org/sqlite"

	"github.com/axseem/blomple/database"
)

func main() {
	sqlite, err := sql.Open("sqlite", "./sqlite.db")
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}
	db := database.New(sqlite)

	c := context.Background()
	article := database.CreateArticleParams{
		Title:   "Hello World",
		Content: "This is the first article",
	}

	if err = db.CreateArticle(c, article); err != nil {
		log.Fatal("failed to create article: ", err)
	}
}
