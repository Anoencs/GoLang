package main

import (
	"crud_app/cli"

	_ "github.com/lib/pq"
)

func main() {
	cli := cli.CommandLine{}
	cli.Run()
}
