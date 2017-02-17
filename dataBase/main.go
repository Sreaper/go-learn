package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbConn, err := sql.Open("mysql", "jmq:jmq@tcp(192.168.179.66:3306)/cap_dev?timeout=90s")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dbConn.Close()
	if rows, err := dbConn.Query("select id ,name ,concat_ws('-',id,name) from application limit 10"); err == nil {
		defer rows.Close()
		for rows.Next() {
			var app Application
			if err = rows.Scan(&app.id, &app.name,&app.combo); err == nil {
				fmt.Println(err)
				fmt.Println(app.id, "-", app.name,"-", app.combo)
			}

		}
	}

}

type Application struct {
	id    int64
	name  string
	combo []byte
}
