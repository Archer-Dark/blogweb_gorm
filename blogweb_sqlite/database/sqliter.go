package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
)

var db  *gorm.DB

func InitSqlite()  {
	fmt.Println("InitSqlite....")
	if db == nil {
		initDb,err := gorm.Open("sqlite3","./database/blogweb.db")
		if err != nil {
			log.Printf("failed to connect database.error: %s",err.Error())
			panic(err)
		}
		db = initDb
	}
}

func GetDB() *gorm.DB {
	return db
}

func init()  {
	InitSqlite()
}