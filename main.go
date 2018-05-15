package main

import (
	l4g "github.com/alecthomas/log4go"
)

func main() {
	l4g.LoadConfiguration("config/log4go.xml")
	l4g.Info("Hello,word!")
	defer l4g.Close()
}
