package main

import (
	"crud_app/cli"
	"crud_app/database"
	"crud_app/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIMPORT_okr_org(t *testing.T) {
	database := database.Database{DbName: "okr", TableName: "okr_org"}
	database.Import_xlsx_okr_org()
	db := database.Db_connect()

	err := db.First(&models.Okr_org{}).Error
	assert.NoError(t, err)

}
func TestIMPORT_okr_period(t *testing.T) {
	database := database.Database{DbName: "okr", TableName: "okr_period"}
	database.Import_xlsx_okr_period()
	db := database.Db_connect()

	err := db.First(&models.Okr_period{}).Error
	assert.NoError(t, err)
}
func TestIMPORT_okr_obj(t *testing.T) {
	database := database.Database{DbName: "okr", TableName: "okr_obj"}
	database.Import_xlsx_okr_obj()
	db := database.Db_connect()

	err := db.First(&models.Okr_obj{}).Error
	assert.NoError(t, err)
}

func TestIMPORT_okr_user(t *testing.T) {
	database := database.Database{DbName: "okr", TableName: "okr_user"}
	database.Import_xlsx_okr_period()
	db := database.Db_connect()

	err := db.First(&models.Okr_user{}).Error
	assert.NoError(t, err)
}

func TestIMPORT_okr_kr(t *testing.T) {
	database := database.Database{DbName: "okr", TableName: "okr_kr"}
	database.Import_xlsx_okr_kr()
	db := database.Db_connect()

	err := db.First(&models.Okr_kr{}).Error
	assert.NoError(t, err)
}

func TestDelete_1(t *testing.T) {
	database := database.Database{DbName: "okr", TableName: "okr_org"}
	db := database.Db_connect()
	database.Delete_by_id("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	if err := db.Where("id = ?", "6ba7b810-9dad-11d1-80b4-00c04fd430c8").First(&models.Okr_org{}).Error; err != nil {
		assert.Error(t, err)
	}
}
func TestDelete_2(t *testing.T) {
	database := database.Database{DbName: "okr", TableName: "okr_org"}
	db := database.Db_connect()
	database.Delete_by_id("12ca82cd-0438-478e-9f0d-910414756a5a")
	if err := db.Where("id = ?", "12ca82cd-0438-478e-9f0d-910414756a5a").First(&models.Okr_org{}).Error; err != nil {
		assert.Error(t, err)
	}
}

func TestDelete_3(t *testing.T) {
	database := database.Database{DbName: "okr", TableName: "okr_org"}
	db := database.Db_connect()
	database.Delete_by_id("c69d4209-44e8-4890-b788-13859c3578ea")
	if err := db.Where("id = ?", "c69d4209-44e8-4890-b788-13859c3578ea").First(&models.Okr_org{}).Error; err != nil {
		assert.Error(t, err)
	}
}

func TestDelete_4(t *testing.T) {
	database := database.Database{DbName: "okr", TableName: "okr_org"}
	db := database.Db_connect()
	database.Delete_by_id("dc00efa3-b2fb-4757-abc4-9c155eacc5a4")
	if err := db.Where("id = ?", "dc00efa3-b2fb-4757-abc4-9c155eacc5a4").First(&models.Okr_org{}).Error; err != nil {
		assert.Error(t, err)
	}
}

func TestDelete_5(t *testing.T) {
	database := database.Database{DbName: "okr", TableName: "okr_org"}
	db := database.Db_connect()
	database.Delete_by_id("b966a15f-1cdd-4187-8dae-3c82fec7f0a7")
	if err := db.Where("id = ?", "b966a15f-1cdd-4187-8dae-3c82fec7f0a7").First(&models.Okr_org{}).Error; err != nil {
		assert.Error(t, err)
	}
}

func TestDelete_6(t *testing.T) {
	database := database.Database{DbName: "okr", TableName: "okr_period"}
	db := database.Db_connect()
	database.Delete_by_id("349d9131-40b2-4513-9bce-e53f98fe5fc8")
	if err := db.Where("id = ?", "349d9131-40b2-4513-9bce-e53f98fe5fc8").First(&models.Okr_period{}).Error; err != nil {
		assert.Error(t, err)
	}
}
func TestDelete_7(t *testing.T) {
	database := database.Database{DbName: "okr", TableName: "okr_period"}
	db := database.Db_connect()
	database.Delete_by_id("35e4b39e-aff2-4377-ab77-e4c99bd6a5fb")
	if err := db.Where("id = ?", "35e4b39e-aff2-4377-ab77-e4c99bd6a5fb").First(&models.Okr_period{}).Error; err != nil {
		assert.Error(t, err)
	}
}

func TestDelete_8(t *testing.T) {
	database := database.Database{DbName: "okr", TableName: "okr_period"}
	db := database.Db_connect()
	database.Delete_by_id("")
	if err := db.Where("id = ?", "c69d4209-44e8-4890-b788-13859c3578ea").First(&models.Okr_period{}).Error; err != nil {
		assert.Error(t, err)
	}
}

func TestDelete_9(t *testing.T) {
	database := database.Database{DbName: "okr", TableName: "okr_period"}
	db := database.Db_connect()
	database.Delete_by_id("3729c0a7-df9b-40d2-bd86-369a344c5e01")
	if err := db.Where("id = ?", "3729c0a7-df9b-40d2-bd86-369a344c5e01").First(&models.Okr_period{}).Error; err != nil {
		assert.Error(t, err)
	}
}

func TestDelete_10(t *testing.T) {
	database := database.Database{DbName: "okr", TableName: "okr_period"}
	db := database.Db_connect()
	database.Delete_by_id("56d2bb58-e63b-4992-86ef-aec86cf442fd")
	if err := db.Where("id = ?", "56d2bb58-e63b-4992-86ef-aec86cf442fd").First(&models.Okr_period{}).Error; err != nil {
		assert.Error(t, err)
	}
}

func TestDelete_11(t *testing.T) {
	database := database.Database{DbName: "okr", TableName: "okr_obj"}
	db := database.Db_connect()
	database.Delete_by_id("021e8321-e796-4a81-8c2b-3eb2e6255079")
	if err := db.Where("id = ?", "021e8321-e796-4a81-8c2b-3eb2e6255079").First(&models.Okr_obj{}).Error; err != nil {
		assert.Error(t, err)
	}
}
func TestDelete_12(t *testing.T) {
	database := database.Database{DbName: "okr", TableName: "okr_obj"}
	db := database.Db_connect()
	database.Delete_by_id("1fe08420-c219-4d02-81ab-fb4315b02733")
	if err := db.Where("id = ?", "1fe08420-c219-4d02-81ab-fb4315b02733").First(&models.Okr_obj{}).Error; err != nil {
		assert.Error(t, err)
	}
}

func TestDelete_13(t *testing.T) {
	database := database.Database{DbName: "okr", TableName: "okr_obj"}
	db := database.Db_connect()
	database.Delete_by_id("2948135b-848d-4219-8889-bd090dbbb285")
	if err := db.Where("id = ?", "2948135b-848d-4219-8889-bd090dbbb285").First(&models.Okr_obj{}).Error; err != nil {
		assert.Error(t, err)
	}
}

func TestDelete_14(t *testing.T) {
	database := database.Database{DbName: "okr", TableName: "okr_obj"}
	db := database.Db_connect()
	database.Delete_by_id("33698075-a535-4fb7-adf6-7e9905c938d0")
	if err := db.Where("id = ?", "33698075-a535-4fb7-adf6-7e9905c938d0").First(&models.Okr_obj{}).Error; err != nil {
		assert.Error(t, err)
	}
}

func TestDelete_15(t *testing.T) {
	database := database.Database{DbName: "okr", TableName: "okr_obj"}
	db := database.Db_connect()
	database.Delete_by_id("5d633179-9f53-41d4-92f6-0c11e13cd4e3")
	if err := db.Where("id = ?", "5d633179-9f53-41d4-92f6-0c11e13cd4e3").First(&models.Okr_obj{}).Error; err != nil {
		assert.Error(t, err)
	}
}

func TestUpdate_1(t *testing.T) {
	cli := cli.CommandLine{}
	okr_org := models.Okr_org{}
	err := okr_org.BeforeUpdate(cli.Update_okr_org("okr", "okr_org", "1535b46e-4c6f-43e8-aeda-d62e1688ce67", "A2"))
	assert.NoError(t, err)

}
func TestUpdate_2(t *testing.T) {
	cli := cli.CommandLine{}
	okr_org := models.Okr_org{}
	err := okr_org.BeforeUpdate(cli.Update_okr_org("okr", "okr_org", "1535b46e-4c6f-43e8-aeda-d62e1688ce67", "A3"))
	assert.NoError(t, err)

}

func TestUpdate_3(t *testing.T) {
	cli := cli.CommandLine{}
	okr_org := models.Okr_org{}
	err := okr_org.BeforeUpdate(cli.Update_okr_org("okr", "okr_org", "1535b46e-4c6f-43e8-aeda-d62e1688ce67", "A4"))
	assert.NoError(t, err)

}

func TestUpdate_4(t *testing.T) {
	cli := cli.CommandLine{}
	okr_org := models.Okr_org{}
	err := okr_org.BeforeUpdate(cli.Update_okr_org("okr", "okr_org", "1535b46e-4c6f-43e8-aeda-d62e1688ce67", "A5"))
	assert.NoError(t, err)

}

func TestUpdate_5(t *testing.T) {
	cli := cli.CommandLine{}
	okr_org := models.Okr_org{}
	err := okr_org.BeforeUpdate(cli.Update_okr_org("okr", "okr_org", "1535b46e-4c6f-43e8-aeda-d62e1688ce67", "A6"))
	assert.NoError(t, err)

}
func TestUpdate_6(t *testing.T) {
	cli := cli.CommandLine{}
	okr_org := models.Okr_org{}
	err := okr_org.BeforeUpdate(cli.Update_okr_org("okr", "okr_org", "243c4557-3dad-403f-9dcb-db4a2b26641f", "H1"))
	assert.NoError(t, err)

}
func TestUpdate_7(t *testing.T) {
	cli := cli.CommandLine{}
	okr_org := models.Okr_org{}
	err := okr_org.BeforeUpdate(cli.Update_okr_org("okr", "okr_org", "243c4557-3dad-403f-9dcb-db4a2b26641f", "H2"))
	assert.NoError(t, err)

}

func TestUpdate_8(t *testing.T) {
	cli := cli.CommandLine{}
	okr_org := models.Okr_org{}
	err := okr_org.BeforeUpdate(cli.Update_okr_org("okr", "okr_org", "243c4557-3dad-403f-9dcb-db4a2b26641f", "H3"))
	assert.NoError(t, err)

}

func TestUpdate_9(t *testing.T) {
	cli := cli.CommandLine{}
	okr_org := models.Okr_org{}
	err := okr_org.BeforeUpdate(cli.Update_okr_org("okr", "okr_org", "243c4557-3dad-403f-9dcb-db4a2b26641f", "H4"))
	assert.NoError(t, err)

}

func TestUpdate_10(t *testing.T) {
	cli := cli.CommandLine{}
	okr_org := models.Okr_org{}
	err := okr_org.BeforeUpdate(cli.Update_okr_org("okr", "okr_org", "243c4557-3dad-403f-9dcb-db4a2b26641f", "H5"))
	assert.NoError(t, err)

}
