package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "modernc.org/sqlite"

	"github.com/axseem/blomple/database"
	"github.com/axseem/blomple/generate"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type HTMLDir struct {
	d http.Dir
}

func (d HTMLDir) Open(name string) (http.File, error) {
	f, err := d.d.Open(name)
	if os.IsNotExist(err) {
		if f, err := d.d.Open(name + ".html"); err == nil {
			return f, nil
		}
	}
	return f, err
}

func ServeSSG() {
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

	if err := generate.IndexPage(rootPath, db); err != nil {
		log.Fatal(err)
	}
	if err := generate.Articles(rootPath, db); err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	index := http.FileServer(HTMLDir{http.Dir(rootPath)})

	r.Handle("/", index)

	fmt.Println("server is running on :1323")
	log.Fatal(http.ListenAndServe(":1323", http.StripPrefix("/", index)))
}
