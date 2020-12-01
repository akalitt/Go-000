package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB

func main() {
	var err error
	DB, err = sql.Open("mysql",
		"root:xxxxx@tcp(127.0.0.1:3306)/testdatabase")

	if err != nil {
		//初始化错误直接panic
		log.Panic(err)
	}

	router := gin.Default()

	router.GET("/user/:id", GetUserByIdController)

	router.Run(":9090")

}
