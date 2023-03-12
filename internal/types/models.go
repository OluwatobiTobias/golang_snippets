package model

import (
	"net/http"

	"github.com/go-chi/render"
)

type Article struct {
	ID     string `json:"id"`
	UserID int16  `json:"user_id"`
	Title  string `json:"title"`
	Slug   string `json:"Slug"`
}

type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging

}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

var ErrNotFound = &ErrResponse{HTTPStatusCode: 404, StatusText: "Resource not found."}
