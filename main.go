package main

import (
	"net/http"
	"os"
	"time"

	log4go "github.com/alecthomas/log4go"

	"io/ioutil"

	"github.com/Unknwon/goconfig"
	//	"github.com/kylelemons/go-gypsy/yaml"
)

func main() {
	log4go.LoadConfiguration("config/log4go.xml")

	var banner string
	if _, err := os.Stat("config/banner.txt"); err != nil {
		banner = "There's no \"config/banner.txt\",You can set it as you like!"
	} else {
		b, err := ioutil.ReadFile("config/banner.txt")
		if err != nil {
			log4go.Error("Read banner err %v ", err)
		}
		banner = string(b)
	}

	var (
		version        string = "1.0.0.0"
		host           string = "127.0.0.1"
		port           string = "8080"
		readtimeout    int    = 10
		writetimeout   int    = 10
		maxheaderbytes int    = 20
	)

	if _, err := os.Stat("config/config.ini"); err != nil {
		log4go.Info("use default config settings")
	} else {
		cfg, err := goconfig.LoadConfigFile("config/config.ini")
		if err != nil {
			log4go.Error("can not load ini file：%s", err)
		}

		version = cfg.MustValue("Server", "version", "1.0.0.0")
		host = cfg.MustValue("Server", "host", "127.0.0.1")
		port = cfg.MustValue("Server", "port", "8080")
		readtimeout = cfg.MustInt("Server", "readtimeout", 10)
		writetimeout = cfg.MustInt("Server", "writetimeout", 10)
		maxheaderbytes = cfg.MustInt("Server", "maxheaderbytes", 20)
	}

	log4go.Info("Server Zgo "+version+" starting at [%s", host+":"+port+"]\n"+banner)

	router := NewRouter(AllRoutes())

	server := &http.Server{
		Addr:           host + ":" + port,
		Handler:        router,
		ReadTimeout:    time.Duration(readtimeout) * time.Second,
		WriteTimeout:   time.Duration(writetimeout) * time.Second,
		MaxHeaderBytes: 1 << uint32(maxheaderbytes),
	}

	log4go.Info(server.ListenAndServe())
}
