package main

import (
	"net/http"
	"time"

	log4go "github.com/alecthomas/log4go"
	"github.com/julienschmidt/httprouter"
)

func Logger(fn func(w http.ResponseWriter, r *http.Request, param httprouter.Params)) func(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
		start := time.Now()
		log4go.Info("%s\t%s	%s -> %s%s\t%s",
			r.Proto,
			r.Method,
			r.RemoteAddr,
			r.Host,
			r.RequestURI,
			r.UserAgent(),
		)

		//		log4go.Info("%s", param)

		fn(w, r, param)
		log4go.Info("Done in %v (%s %s)", time.Since(start), r.Method, r.URL.Path)
	}
}
