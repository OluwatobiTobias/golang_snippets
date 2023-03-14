package handlers

import (
	"fmt"
	"net/http"

	er "github.com/OluwatobiTobias/golang_snippets/internal/error"
	middleware "github.com/OluwatobiTobias/golang_snippets/internal/middleware"
	res "github.com/OluwatobiTobias/golang_snippets/internal/resp_req"
	types "github.com/OluwatobiTobias/golang_snippets/internal/types"
	"github.com/go-chi/render"
)

// GetArticle returns the specific Article. You'll notice it just
// fetches the Article right off the context, as its understood that
// if we made it this far, the Article must be on the context. In case
// its not due to a bug, then it will panic, and our Recoverer will save us.
func GetArticle(w http.ResponseWriter, r *http.Request) {
	// Assume if we've reach this far, we can access the article
	// context because this handler is a child of the ArticleCtx
	// middleware. The worst case, the recoverer middleware will save us.
	a := r.Context().Value(middleware.ConKey)
	fmt.Println("article----", a)
	article := r.Context().Value(middleware.ConKey).(*types.Article)

	if err := render.Render(w, r, res.NewArticleResponse(article)); err != nil {
		render.Render(w, r, er.ErrRender(err))
		return
	}
}
