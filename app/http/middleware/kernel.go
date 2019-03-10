package middleware

import "github.com/kataras/iris/context"

// Middlewares table
func Middlewares() []context.Handler {

	return []context.Handler{
		HideID,
	}
}
