package contracts

import "net/http"

type Middleware interface {
	Handle(w http.ResponseWriter, r *http.Request) bool
}
