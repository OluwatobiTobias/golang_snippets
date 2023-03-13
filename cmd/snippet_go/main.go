package main

import (
	"fmt"
	"net/http"

	handler "github.com/OluwatobiTobias/golang_snippets/internal/handlers"
	custom_middle "github.com/OluwatobiTobias/golang_snippets/internal/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("home sweet"))
	})

	r.Route("/{articleID}", func(r chi.Router) {
		fmt.Println("go here")
		r.Use(custom_middle.ArticleCtx)
		r.Get("/", handler.GetArticle)
	})

	http.ListenAndServe(":3333", r)
}
