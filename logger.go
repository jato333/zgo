package main

import (
	"net/http"
	"time"

	l4g "github.com/alecthomas/log4go"
)

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		l4g.Info(
			"%s\t%s\t%s -> %s%s\t%s\t%s",
			r.Proto,
			r.Method,
			r.RemoteAddr,
			r.Host,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
