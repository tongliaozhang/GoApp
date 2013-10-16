package db

import (
	"GoApp/fileDeal"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
)

func InsertMysql(data *fileDeal.DataFile, commond string) {

	db, _ := sql.Open("mysql", "root:zhangql@tcp(localhost:3306)/mysql_study")
	defer db.Close()

	datas := data.Data

	coon, _ := db.Begin()

	fmt.Println(len(datas))

	for i := 0; i < len(datas)-1; i++ {
		tmp := datas[i]
		args := strings.Split(tmp, ",")

		_, err := db.Exec(commond, args[0], args[1], args[2], args[3])

		if err != nil {
			log.Fatal(err)
		}
	}

	coon.Commit()

}
