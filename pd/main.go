package main

import (
	"crud_app/cli"

	_ "github.com/lib/pq"
)

func main() {
	cli := cli.CommandLine{}
	cli.Run()

	// xlsx := xlsx.Xlsx{FilePath: "OKR.xlsx", SheetName: "OKR"}
	// okr_period := models.Okr_period{}
	// okr_user := models.Okr_user{}
	// okr_org := models.Okr_org{}
	// obj := models.Okr_obj{}

	// okr_obj, okr_kr := obj.Read(xlsx)
	// okr_user.Read(xlsx)
	// okr_org.Read(xlsx)
	// okr_period.Read(xlsx)
	// okr_user.Org_id = okr_org.Id
	// okr_org.Okr_user = okr_user

	// for i := 0; i < len(okr_obj); i++ {
	// 	okr_kr[i].Create_by = okr_user.User_id
	// }
	// spew.Dump(okr_kr[2])

}
