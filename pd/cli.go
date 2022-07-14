package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
)

type CommandLine struct{}

func (cli *CommandLine) printUsage() {
	fmt.Println("Usage:")
	fmt.Println("delete -dbname DBNAME -tbname TBNAME -id ID -delete by id")
	fmt.Println("import -dbname DBNAME -tbname TBNAME")
	fmt.Println("read -dbname DBNAME -tbname TBNAME")

}

func (cli *CommandLine) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		runtime.Goexit()
	}
}

func (cli *CommandLine) Run() {
	cli.validateArgs()
	// delete flag
	deleteByIdCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	delete_ID := deleteByIdCmd.String("id", "", "The id to delete")
	delete_dbname := deleteByIdCmd.String("dbname", "", "The db name to delete")
	delete_tablename := deleteByIdCmd.String("tbname", "", "The table name to delete")
	// list flag
	listCmd := flag.NewFlagSet("read", flag.ExitOnError)
	list_dbname := listCmd.String("dbname", "", "The db name to import")
	list_tbname := listCmd.String("tbname", "", "The tb name to import")
	//import flag
	importCmd := flag.NewFlagSet("import", flag.ExitOnError)
	import_dbname := importCmd.String("dbname", "", "The db name to import")
	import_tbname := importCmd.String("tbname", "", "The tb name to import")
	// update flag
	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	update_dbname := updateCmd.String("dbname", "", "The db name to import")
	update_tbname := updateCmd.String("tbname", "", "The tb name to import")                                                                                        
	update_id := updateCmd.String("id", "", "The id to update")
	update_name := updateCmd.String("name", "", "The name data")
	update_month := updateCmd.String("month", "", "The month data")
	update_year := updateCmd.String("year", "", "The year data")
	update_quarter := updateCmd.String("quarter", "", "The quarter data")

	switch os.Args[1] {
	case "delete":
		err := deleteByIdCmd.Parse(os.Args[2:])
		if err != nil {
			log.Fatal(err)
		}
	case "import":
		err := importCmd.Parse(os.Args[2:])
		if err != nil {
			log.Fatal(err)
		}
	case "list":
		err := listCmd.Parse(os.Args[2:])
		if err != nil {
			log.Fatal(err)
		}
	case "update":
		err := updateCmd.Parse(os.Args[2:])
		if err != nil {
			log.Fatal(err)
		}
	}

	if deleteByIdCmd.Parsed() {
		if *delete_ID == "" || *delete_dbname == "" || *delete_tablename == "" {
			deleteByIdCmd.Usage()
			runtime.Goexit()
		}
		cli.delete(*delete_dbname, *delete_tablename, *delete_ID)
	}

	if importCmd.Parsed() {
		if *import_dbname == "" || *import_tbname == "" {
			importCmd.Usage()
			runtime.Goexit()
		}
		cli.import_xlsx(*import_dbname, *import_tbname)
	}

	if listCmd.Parsed() {
		if *list_dbname == "" || *list_tbname == "" {
			importCmd.Usage()
			runtime.Goexit()
		}
		cli.read(*list_dbname, *list_tbname)
	}
	if updateCmd.Parsed() {
		if *update_dbname == "" || *update_tbname == "" {
			importCmd.Usage()
			runtime.Goexit()
		}
		if *update_tbname == "okr_org" {
			cli.update_okr_org(*update_dbname, *update_tbname, *update_id, *update_name)
		} else if *update_tbname == "okr_period" {
			cli.update_okr_period(*update_dbname, *update_tbname, *update_id, *update_month, *update_year, *update_quarter, *update_name)
		}
	}
}
