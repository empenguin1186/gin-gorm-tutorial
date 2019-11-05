package main

import (
	"github.com/cihub/seelog"
	"github.com/empenguin1186/gin-gorm-tutorial/server"
)

func main() {

	// ロガー設定
	logger, err := seelog.LoggerFromConfigAsFile("seelog.xml")
	if err != nil {
		panic(err)
	}
	seelog.ReplaceLogger(logger)
	// db.Init()
	server.Init()
	// db.Close()
}
