package main

import (
	"time"

	log4go "github.com/alecthomas/log4go"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(routes Routes) *httprouter.Router {
	router := httprouter.New()
	log4go.Info("Begin Load service ...")
	start := time.Now()
	for _, route := range routes {
		var handle httprouter.Handle

		handle = route.HandlerFunc
		handle = Logger(handle)

		router.Handle(route.Method, route.Path, handle)
		log4go.Info("Loading [%s]\t%s", route.Method, route.Path)
	}

	log4go.Info("%d services were loaded successfully! Total cost %s", len(routes), time.Since(start))

	return router
}
