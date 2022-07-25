package main

import (
	"crud_app/models"

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
	// config, err := util.LoadConfig(".")
	// if err != nil {
	// 	log.Fatal("cannot load config:", err)
	// }
	// LOG_FILE := config.LOG_FILE
	// logFile, err := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	// if err != nil {
	// 	log.Panic(err)
	// 	return
	// }
	// defer logFile.Close()
	// log.SetOutput(logFile)
	// cli := cli.CommandLine{}
	// cli.Run()

	// fileLV := []fileLevel{}
	// _ = filepath.Walk("ITD", func(path string, info os.FileInfo, err error) error {
	// 	//fmt.Println(filepath.Base(filepath.Dir(path)))
	// 	// if !slices.Contains(files, filepath.Base(filepath.Dir(path))) {
	// 	// 	files = append(files, filepath.Base(filepath.Dir(path)))
	// 	// }
	// 	filelv := fileLevel{}
	// 	filelv.Name = filepath.Base(filepath.Dir(path))
	// 	filelv.Level = len(strings.Split(path, "\\")) - 1
	// 	if !slices.Contains(fileLV, filelv) {
	// 		if filepath.Base(filepath.Dir(path)) != "." {
	// 			//files = append(files, filepath.Base(filepath.Dir(path)))
	// 			fileLV = append(fileLV, filelv)
	// 		}

	// 	}
	// 	//	files = append(files, filepath.Base(filepath.Dir(path)))
	// 	if err != nil {
	// 		log.Panic(err)
	// 	}
	// 	return nil
	// })
	// //fmt.Println(len(files))
	// fmt.Println(len(fileLV))
	// for _, item := range fileLV {
	// 	spew.Dump(item)
	// }

	org := models.Okr_org{}
	org.Read()
}
