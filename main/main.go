package main

import (
	"database/sql"
	"fmt"

	"github.com/ridwanfathin/blog-server/router"

	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	fmt.Println("Welcome to the webserver")

	e := router.New()

	e.Start(":8000")

}
