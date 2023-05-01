package main

import (
	"io/ioutil"
	"log"

	"github.com/fatram/golang-dansmultipro-test/config"
	"github.com/fatram/golang-dansmultipro-test/internal/connector"
)

func main() {
	config.ReadConfig("./.env")
	db := connector.LoadMysqlDatabase()
	c, ioErr := ioutil.ReadFile("./migration/init.sql")
	if ioErr != nil {
		log.Panic("file does not exist")
	}
	sql := string(c)
	_, err := db.Exec(sql)
	if err != nil {
		log.Panic(err)
	}
}
