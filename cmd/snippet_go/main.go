package main

import (
	"flag"

	custom_middleware "github.com/OluwatobiTobias/golang_snippets/internal/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

// var routes = flag.Bool("routedoc", false, "Generate router docs")

func main() {
	flag.Parse()

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Route("/{articles}", func(r chi.Router) {
		r.Use(custom_middleware.ArticleCtx)
		// r.Get("/", GetArticle)
		// 	r.Put("/", UpdateArticle)
		// 	r.Delete("/", DeleteArticle)
	})

}
