package main

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/Unknwon/goconfig"
	l4g "github.com/alecthomas/log4go"
)

func main() {
	//加载日志配置文件
	l4g.LoadConfiguration("config/log4go.xml")
	defer l4g.Close()

	//加载配置文件
	cfg, err := goconfig.LoadConfigFile("config/config.ini")
	if err != nil {
		l4g.Error("读取配置文件失败[config/config.ini]")
		return
	}

	//读取配置字段信息（每个字段都给出响应的默认值）
	host := cfg.MustValue("Server", "host", "127.0.0.1")
	port := cfg.MustValue("Server", "port", "8080")
	version := cfg.MustValue("Server", "version", "1.0.0")
	readtimeout := cfg.MustInt("Server", "readtimeout", 10)
	writetimeout := cfg.MustInt("Server", "writetimeout", 10)
	maxheaderbytes := cfg.MustInt("Server", "maxheaderbytes", 10)

	//加载banner
	b, err := ioutil.ReadFile("config/banner.txt")
	if err != nil {
		l4g.Error("读取banner文件失败[config/banner.txt]")
	}
	banner := string(b)

	l4g.Info("Server Zgo (" + version + ") starting at:[" + host + ":" + port + "]  ...\n" + banner)
	router := NewRouter()
	//设置服务端信息
	server := &http.Server{
		Addr:           host + ":" + port,
		Handler:        router,
		ReadTimeout:    time.Duration(readtimeout) * time.Second,
		WriteTimeout:   time.Duration(writetimeout) * time.Second,
		MaxHeaderBytes: 1 << uint16(maxheaderbytes),
	}

	l4g.Info(server.ListenAndServe())
}
