package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	PORT      = 0
	SECRETKEY []byte
	DBDATAURL  = ""
	DBNAME  = ""
	DBDRIVER  = ""
	DBURL     = ""
)

// Load the server PORT
func init() {
	var err error
	err = godotenv.Load("config.env")
	if err != nil {
		log.Fatal(err)
	}
	PORT, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		PORT = 9000
	}

	DBDATAURL=fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/", os.Getenv("DB_USER"), os.Getenv("DB_PassWord"), )
	DBNAME = os.Getenv("DB_NAME")
	DBDRIVER = os.Getenv("DB_DRIVER")
	DBURL = fmt.Sprintf("%s:%s@/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PassWord"),
		os.Getenv("DB_NAME"),
	) + "?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"
	SECRETKEY = []byte(os.Getenv("API_SECRET"))
}