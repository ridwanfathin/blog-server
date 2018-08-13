package postgres

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//OpenDB open conncection to postgresql database
func OpenDB() *gorm.DB {
	db, err := gorm.Open("postgres",
		`host=localhost
		user=username password=password
		dbname=databasename
		sslmode=disable`)
	if err != nil {
		panic(err)
	}

	return db
}
