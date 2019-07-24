package main

import (
	_ "github.com/go-sql-driver/mysql"

	"database/sql"
	"fmt"
	"github.com/bagusandrian/goclass/dbconnection/src/config"
	"log"
)

var conf *config.Config

func init() {
	conf = config.ReadConfig()
	log.Printf("%+v\n", conf)
}
func main() {
	db, err := sql.Open("mysql", conf.Database.DBMaster)
	checkErr(err)

	// insert
	stmt, err := db.Prepare("insert into `userinfo` (username, department, created) values (?, ?, ?)")
	checkErr(err)

	res, err := stmt.Exec("admin", "Administrator", "2019-01-01")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)

	// update
	stmt, err = db.Prepare("update `userinfo` set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("adminupdate", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	// query
	rows, err := db.Query("select * from userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	// delete
	stmt, err = db.Prepare("delete from `userinfo` where uid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
