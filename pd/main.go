package main

import (
	"crud_app/cli"

	_ "github.com/lib/pq"
)

func main() {
	cli := cli.CommandLine{}
	cli.Run()

	// err := filepath.Walk("ITD", func(path string, info os.FileInfo, err error) error {
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return err
	// 	}
	// 	fmt.Printf("dir: %v: name: %s\n", info.IsDir(), path)
	// 	return nil
	// })
	// if err != nil {
	// 	fmt.Println(err)
	// }

}
