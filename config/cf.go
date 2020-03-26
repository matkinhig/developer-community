package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	PORT      = 0
	SECRETKEY []byte
	DBDRIVE   = ""
	DBURL     = ""
)

func Load() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	PORT, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		PORT = 3456
	}
	DBDRIVE = os.Getenv("DB_DRIVER")
	DBURL = fmt.Sprintf("%s:%s@(127.0.0.1)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))
	SECRETKEY = []byte(os.Getenv("API_SECRET"))
}
