package main

import (
	"net/http"

	l4g "github.com/alecthomas/log4go"
)

func main() {
	l4g.LoadConfiguration("config/log4go.xml")
	defer l4g.Close()

	router := NewRouter()

	l4g.Info(http.ListenAndServe(":8080", router))
}
