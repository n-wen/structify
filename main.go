package main

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/alecthomas/kingpin"
	_ "github.com/go-sql-driver/mysql"
)

var tag = "db"

func main() {
	password := kingpin.Flag("password", "password").Short('p').String()
	username := kingpin.Flag("username", "username").Short('u').String()
	host := kingpin.Flag("host", "host").Short('h').Default("127.0.0.1").String()
	port := kingpin.Flag("port", "port").Short('P').Default("3306").Int()
	database := kingpin.Flag("db", "database").Short('d').String()
	table := kingpin.Flag("table", "table").Short('t').String()
	kingpin.Parse()
	if username == nil || database == nil || table == nil {
		panic("missing params")
	}
	var (
		USERNAME = *username
		PASSWORD = *password
		NETWORK  = "tcp"
		SERVER   = *host
		PORT     = *port
		DATABASE = *database
	)
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=UTC", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	rows, err := db.Query("show full columns from " + *table)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	type Col struct {
		Field      string
		Type       string
		Collation  interface{}
		Null       interface{}
		Key        interface{}
		Default    interface{}
		Extra      interface{}
		Privileges interface{}
		Comment    string
	}
	var cols []Col
	var p PrintAtom
	p.Add("type", BigCamelMarshal(*table), "struct {")
	for rows.Next() {
		var col Col
		err := rows.Scan(&col.Field, &col.Type, &col.Collation, &col.Null, &col.Key, &col.Default, &col.Extra, &col.Privileges, &col.Comment)
		if err != nil {
			panic(err.Error())
		}
		cols = append(cols, col)
		p.Add(BigCamelMarshal(col.Field), getTypeName(col.Type, false), fmt.Sprintf("`%v:\"%v\"`", tag, col.Field),
			fmt.Sprintf(" // %v", col.Comment))
	}
	p.Add("}")
	err = rows.Err()
	if err != nil {
		panic(err.Error())
	}
	output := strings.Join(p.Generates(), "\n")
	fmt.Printf("%v\n", output)
}
