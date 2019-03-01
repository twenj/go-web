package go_web

import (
	"fmt"
	"go-web/utils"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type App struct {
	AppPath string
	ControllerPath string
	ConfigPath string
	Conf *Config
  Server *http.Server
}

func (app *App) Run() {
	// 获取 app 目录
	app.initPath()
	// 配置信息初始化
	app.Conf = ConfigInit(app.ConfigPath)
	r := RouterBind(app.ControllerPath)
	
	http.HandleFunc("/", r)
	fmt.Println("Server is running...")
	err := http.ListenAndServe(":7140", nil)
	if err != nil {
		log.Fatal("ListenAndServer: " + err.Error())
	}
}

func (app *App) initPath() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	Error(err)
	apppath := utils.Substr(dir, 0, -5)
	app.AppPath = apppath
	app.ConfigPath = apppath + "config/"
	app.ControllerPath = apppath + "controllers/"
}

