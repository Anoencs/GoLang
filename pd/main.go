package main

import (
	"crud_app/cli"
	"crud_app/util"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// type fileLevel struct {
// 	Name  string
// 	Level int
// }

// func (item1 fileLevel) compare(item2 fileLevel) bool {
// 	if item1.Name == item2.Name {
// 		return true
// 	}
// 	return false
// }
func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	LOG_FILE := config.LOG_FILE
	logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
		return
	}
	defer logFile.Close()
	log.SetOutput(logFile)
	cli := cli.CommandLine{}
	cli.Run()
}
