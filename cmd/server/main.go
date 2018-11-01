package main

import (
	"flag"
	"net/http"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	address := flag.String("p", ":8080", "Address where server will listen to.")
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK."))
	})
	
	http.ListenAndServe(*address, r)
}
