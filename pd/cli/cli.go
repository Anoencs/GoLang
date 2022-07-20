package cli

import (
	"crud_app/database"
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
	// init db flag
	initCmd := flag.NewFlagSet("init", flag.ContinueOnError)
	// delete flag
	deleteByIdCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	delete_ID := deleteByIdCmd.String("id", "", "The id to delete")
	delete_dbname := deleteByIdCmd.String("dbname", "", "The db name to delete")
	delete_tablename := deleteByIdCmd.String("tbname", "", "The table name to delete")

	importCmd := flag.NewFlagSet("import", flag.ExitOnError)
	import_xlsx := importCmd.String("xlsx", "", "The excel path")
	import_sheet := importCmd.String("sheet", "", "The sheet name")
	import_folder := importCmd.String("dir", "", "The dir")
	// update flag
	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	update_dbname := updateCmd.String("dbname", "", "The db name to import")
	update_tbname := updateCmd.String("tbname", "", "The tb name to import")
	update_id := updateCmd.String("id", "", "The id to update")
	update_name := updateCmd.String("name", "", "The name data")
	update_month := updateCmd.Uint64("month", 0, "The month data")
	update_year := updateCmd.Uint64("year", 0, "The year data")
	update_quarter := updateCmd.Uint64("quarter", 0, "The quarter data")
	update_itype := updateCmd.String("itype", "", "The quarter data")
	update_criterias := updateCmd.Uint64("criterias", 0, "The quarter data")
	update_start := updateCmd.Float64("start", 0, "The quarter data")
	update_target := updateCmd.String("target", "", "The quarter data")
	update_selfgrade := updateCmd.Float64("selfgrade", 0, "The quarter data")
	update_grade := updateCmd.Float64("grade", 0, "The quarter data")
	update_org_id := updateCmd.String("orgid", "", "The org id data")
	update_obj_id := updateCmd.String("objid", "", "The org id data")
	update_user_id := updateCmd.String("userid", "", "The org id data")
	update_period_id := updateCmd.String("periodid", "", "The org id data")
	update_manager_id := updateCmd.String("mngid", "", "The org id data")

	update_email := updateCmd.String("email", "", "The org id data")
	update_manager_email := updateCmd.String("mngemail", "", "The org id data")
	update_role := updateCmd.String("role", "", "The org id data")
	update_department := updateCmd.String("department", "", "The org id data")

	update_review_date := updateCmd.String("rvdate", "", "The org id data")
	update_create_date := updateCmd.String("createdate", "", "The org id data")
	update_create_by := updateCmd.String("createby", "", "The org id data")
	update_last_modified := updateCmd.String("lastmodifi", "", "The org id data")
	update_last_modified_by := updateCmd.String("lastmodifiby", "", "The org id data")
	update_status := updateCmd.Uint64("status", 0, "status")
	update_create := updateCmd.String("create", "", "The org id data")
	update_targetdate := updateCmd.String("targetdate", "", "The org id data")
	update_duedate := updateCmd.String("duedate", "", "The org id data")
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
	case "update":
		err := updateCmd.Parse(os.Args[2:])
		if err != nil {
			log.Fatal(err)
		}
	case "init":
		err := initCmd.Parse(os.Args[2:])
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
		if *import_xlsx == "" && *import_folder != "" {
			cli.import_all_xlsx_folder(*import_folder)
		} else if *import_xlsx == "" && *import_folder == "" {
			importCmd.Usage()
			runtime.Goexit()
		}
		if *import_sheet == "" && *import_xlsx != "" {
			cli.import_all_xlsx(*import_xlsx)
		} else if *import_sheet != "" && *import_xlsx != "" {
			cli.import_xlsx(*import_xlsx, *import_sheet)
		}
	}

	if updateCmd.Parsed() {
		if *update_dbname == "" || *update_tbname == "" || *update_id == "" {
			updateCmd.Usage()
			runtime.Goexit()
		}
		if *update_tbname == "okr_periods" {
			cli.Update_okr_period(*update_dbname, *update_tbname, *update_id, *update_name, *update_month, *update_year, *update_quarter)
		} else if *update_tbname == "okr_orgs" {
			cli.Update_okr_org(*update_dbname, *update_tbname, *update_id, *update_name)
		} else if *update_tbname == "okr_objs" {
			cli.Update_okr_obj(*update_dbname, *update_tbname, *update_id, *update_name, *update_org_id, *update_user_id, *update_period_id, *update_review_date, *update_create_date, *update_create_by, *update_last_modified, *update_last_modified_by, *update_status)
		} else if *update_tbname == "okr_users" {
			cli.Update_okr_user(*update_dbname, *update_tbname, *update_org_id, *update_name, *update_user_id, *update_manager_id, *update_email, *update_manager_email, *update_role, *update_department)
		} else {
			cli.Update_okr_kr(*update_dbname, *update_tbname, *update_create, *update_last_modified, *update_duedate, *update_obj_id, *update_name, *update_id, *update_user_id, *update_last_modified_by, *update_targetdate, *update_create_by, *update_itype, *update_target, *update_criterias, *update_start, *update_selfgrade, *update_grade)
		}
	}
	if initCmd.Parsed() {
		db := database.Database{DbName: "okr"}
		db.Init()
	}
}
