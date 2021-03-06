package main

import (
	"assets"
	"database/sql"
	"fmt"
	"github.com/go-zoo/bone"
	_ "github.com/lib/pq"
	"github.com/unrolled/render"
	"log"
	"net/http"
)

func main() {
	db, err := sql.Open("postgres", "user=laptop dbname=estelle_test sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	r := render.New(render.Options{
		Directory:  "views",
		Extensions: []string{".html"},
	})

	mux := bone.New()
	ServeResource := assets.ServeResource
	mux.HandleFunc("/img/", ServeResource)
	mux.HandleFunc("/css/", ServeResource)
	mux.HandleFunc("/js/", ServeResource)

	mux.HandleFunc("/pages", func(w http.ResponseWriter, req *http.Request) {
		rows, err := db.Query("SELECT id, title FROM pages")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		type yourtype struct {
			Id    int
			Title string
		}

		s := []yourtype{}
		for rows.Next() {
			var t yourtype
			if err := rows.Scan(&t.Id, &t.Title); err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s", t.Title)
			s = append(s, t)
		}
		r.HTML(w, http.StatusOK, "foofoo", s)
	})

	mux.HandleFunc("/bar", func(w http.ResponseWriter, req *http.Request) {
		r.HTML(w, http.StatusOK, "bar", nil)
	})

	mux.HandleFunc("/home/:id", func(w http.ResponseWriter, req *http.Request) {
		id := bone.GetValue(req, "id")
		r.HTML(w, http.StatusOK, "index", id)
	})

	mux.HandleFunc("/foo", func(w http.ResponseWriter, req *http.Request) {
		r.HTML(w, http.StatusOK, "foo", nil)
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		r.HTML(w, http.StatusOK, "index", nil)
	})

	http.ListenAndServe(":8080", mux)
}
