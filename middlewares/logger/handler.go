package logger

import (
	"fmt"
	"net/http"
)

type loggerMiddleware struct{}

func Handler() *loggerMiddleware {
	return &loggerMiddleware{}
}

func (_ *loggerMiddleware) Handle(w http.ResponseWriter, r *http.Request) bool {
	fmt.Printf("%s %s → %s %s%s\n", r.RemoteAddr, r.Header["User-Agent"], r.Method, r.Host, r.URL)
	return true
}
