package auto

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"gotool/api/database"
	"gotool/api/models"
	"gotool/config"
	"log"
)

func init() {
	db, err := sql.Open(config.DBDRIVER, config.DBDATAURL)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	_, err = db.Exec("CREATE DATABASE " + config.DBNAME)
	if err != nil {
		log.Println("数据库已存在!")
		InitDatabase()
		return
	}
	log.Println("数据库创建成功！", err)
	InitDatabase()
}

func InitDatabase() {
	err := database.InitDb()
	if err != nil {
		log.Fatal("Gorm初始化数据库失败！报错：" + err.Error())
	}
}

func Load() {
	db := database.NewDb()
	err := db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatal(err)
	}
}
