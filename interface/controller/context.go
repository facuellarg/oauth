package controller

import "context"

type Context interface {
	Bind(i interface{}) error
	Query(key string) string
	Context() context.Context
}
