package main

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
	"log"

	//github lib
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jackc/pgx/v5/pgxpool"

	//connect db packet
	"snip/storage"
)

type PageData struct {
	ShortURL string
}

func main() {
	ctx := context.Background()
	
	connStr := os.Getenv("DB_CONN")
	if connStr == "" {
		connStr = "postgres://postgres:password@localhost:5432/snip_db?sslmode=disable"
	}

	dbPool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		fmt.Printf("main.go error: error connect DB : %w",err)
		os.Exit(1)
	}

	defer dbPool.Close()

	//initialising DB
	store := storage.New(dbPool)

	//upload html
	tmplHTML := template.Must(template.ParseFiles("index.html"))

	//router
	router := chi.NewRouter()

	//log
	router.Use(middleware.DefaultLogger)
	//panic killer
	router.Use(middleware.Recoverer)

	//static
	staticFile := http.FileServer(http.Dir("./static"))
	router.Handle("/static/*",http.StripPrefix("/static/",staticFile))
	//
	
	//map
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		tmplHTML.Execute(w,nil)
	})

	router.Post("/",func(w http.ResponseWriter, r *http.Request) {
		ClassicURL := r.FormValue("url")

		if ClassicURL == "" {
			http.Redirect(w,r,"/",http.StatusSeeOther)
			return
		}

		hash,err := SnipHashAlgoritm(ClassicURL)
		if err != nil {
			http.Error(w,"error create link",http.StatusInternalServerError)
			return
		}

		HashShortURL := "http://" + r.Host + "/" +  hash

		err =  store.SaveUrl(r.Context(), ClassicURL,hash) 
		if err != nil {
			log.Printf("DB Error: %v\n", err)
			http.Error(w,"error saving DB",http.StatusInternalServerError)
			return
		}
		data := PageData{ShortURL: HashShortURL}
		tmplHTML.Execute(w,data)
	})

	router.Get("/{hash}",func(w http.ResponseWriter, r *http.Request) {
		hash := chi.URLParam(r,"hash")
		originalUrl, err := store.GetUrl(r.Context(),hash)
		if err != nil {
			http.NotFound(w,r)
			return
		}
		if !strings.HasPrefix(originalUrl, "http://") && !strings.HasPrefix(originalUrl, "https://"){
			originalUrl = "https://" + originalUrl
		}

		http.Redirect(w,r,originalUrl,http.StatusFound)
	})

	fmt.Printf("Server Active. Port: 8080")
	http.ListenAndServe(":8080",router)
}