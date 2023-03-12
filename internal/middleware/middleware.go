package middleware

import (
	"context"
	"net/http"

	query "github.com/OluwatobiTobias/golang_snippets/internal/database"
	model "github.com/OluwatobiTobias/golang_snippets/internal/types"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func ArticleCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var article *model.Article
		var err error

		if articleID := chi.URLParam(r, "articleID"); articleID != "" {
			article, err = query.DbGetArticleById(articleID)
		} else if articleSlug := chi.URLParam(r, "articleSlug"); articleSlug != "" {
			article, err = query.DbGetArticleBySlug(articleSlug)
		} else {
			render.Render(w, r, model.ErrNotFound)
		}

		if err != nil {
			render.Render(w, r, model.ErrNotFound)
			return
		}

		ctx := context.WithValue(r.Context(), contextKeyRequestID, article)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

type contextKey int

const (
	contextKeyRequestID contextKey = iota
)
