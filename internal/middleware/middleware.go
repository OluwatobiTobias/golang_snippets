package middleware

import (
	"context"
	"fmt"
	"net/http"

	database "github.com/OluwatobiTobias/golang_snippets/internal/database"
	er "github.com/OluwatobiTobias/golang_snippets/internal/error"
	types "github.com/OluwatobiTobias/golang_snippets/internal/types"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

// ArticleCtx middleware is used to load an Article object from
// the URL parameters passed through as the request. In case
// the Article could not be found, we stop here and return a 404.
func ArticleCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var article *types.Article
		var err error

		fmt.Println("got to middle func")

		if articleID := chi.URLParam(r, "articleID"); articleID != "" {
			article, err = database.DBGetArticle(articleID)
		} else if articleSlug := chi.URLParam(r, "articleSlug"); articleSlug != "" {
			article, err = database.DBGetArticleBySlug(articleSlug)
		} else {
			fmt.Println("render else error instead")
			render.Render(w, r, er.ErrNotFound)
			return
		}
		if err != nil {
			render.Render(w, r, er.ErrNotFound)
			fmt.Println("render err error instead")
			return
		}
		fmt.Printf("%v---%T\n", article, article)
		ctx := context.WithValue(r.Context(), ConKey, article)
		fmt.Println("ctx was------", ctx.Value(ConKey))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

type contextKey int

const ConKey contextKey = iota + 1
