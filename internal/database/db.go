package database

import (
	"errors"
	"fmt"
	"math/rand"

	types "github.com/OluwatobiTobias/golang_snippets/internal/types"
)

// Article fixture data
var articles = []*types.Article{
	{ID: "1", UserID: 100, Title: "Hi", Slug: "hi"},
	{ID: "2", UserID: 200, Title: "sup", Slug: "sup"},
	{ID: "3", UserID: 300, Title: "alo", Slug: "alo"},
	{ID: "4", UserID: 400, Title: "bonjour", Slug: "bonjour"},
	{ID: "5", UserID: 500, Title: "whats up", Slug: "whats-up"},
}

// User fixture data
var users = []*types.User{
	{ID: 100, Name: "Peter"},
	{ID: 200, Name: "Julia"},
}

func DBNewArticle(article *types.Article) (string, error) {
	article.ID = fmt.Sprintf("%d", rand.Intn(100)+10)
	articles = append(articles, article)
	return article.ID, nil
}

func DBGetArticle(id string) (*types.Article, error) {
	for _, a := range articles {
		if a.ID == id {
			return a, nil
		}
	}
	return nil, errors.New("article not found")
}

func DBGetArticleBySlug(slug string) (*types.Article, error) {
	for _, a := range articles {
		if a.Slug == slug {
			return a, nil
		}
	}
	return nil, errors.New("article not found")
}

func DBUpdateArticle(id string, article *types.Article) (*types.Article, error) {
	for i, a := range articles {
		if a.ID == id {
			articles[i] = article
			return article, nil
		}
	}
	return nil, errors.New("article not found")
}

func DBRemoveArticle(id string) (*types.Article, error) {
	for i, a := range articles {
		if a.ID == id {
			articles = append((articles)[:i], (articles)[i+1:]...)
			return a, nil
		}
	}
	return nil, errors.New("article not found")
}

func DBGetUser(id int64) (*types.User, error) {
	for _, u := range users {
		if u.ID == id {
			return u, nil
		}
	}
	return nil, errors.New("user not found")
}