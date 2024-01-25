package main

import (
	"context"
	"database/sql"
	"log"
	"os"

	_ "embed"

	_ "modernc.org/sqlite"

	"github.com/axseem/blomple/database"
)

func main() {
	sqlite, err := sql.Open("sqlite", "./sqlite.db")
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}
	db := database.New(sqlite)

	md, err := os.ReadFile("./cmd/seed/features.md")
	if err != nil {
		log.Fatal(err)
	}

	article := database.CreateArticleParams{
		Title:   "Markdown features showcase",
		Content: string(md),
	}

	if err = db.CreateArticle(context.Background(), article); err != nil {
		log.Fatal("failed to create article: ", err)
	}
}
