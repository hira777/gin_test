package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// DBConnection is a database handle representing a pool of zero or more underlying connections.
var DBConnection *sql.DB

// Person Structt
type Person struct {
	Name string
	Age  int
}

func main() {
	DBConnection, _ := sql.Open("sqlite3", "./example.sql")
	defer DBConnection.Close()
	// person テーブルが存在しなければ作成する
	cmd := `CREATE TABLE IF NOT EXISTS person(
		name STRING,
		age INT)`
	_, err := DBConnection.Exec(cmd)
	if err != nil {
		log.Fatal(err)
	}

	// データの追加
	// cmd = "INSERT INTO person (name, age) VALUES (?, ?)"
	// _, err = DBConnection.Exec(cmd, "Nancy", 20)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// データの更新
	// cmd = "UPDATE person SET age = ? WHERE name = ?"
	// _, err = DBConnection.Exec(cmd, 25, "Mike")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// データの取得
	// cmd = "SELECT * from person"
	// rows, _ := DBConnection.Query(cmd)
	// defer rows.Close()
	// var pp []Person
	// for rows.Next() {
	// 	var p Person
	// 	err := rows.Scan(&p.Name, &p.Age)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	pp = append(pp, p)
	// }
	// for _, p := range pp {
	// 	fmt.Println(p.Name, p.Age)
	// }

	// データの取得（条件付き）
	// cmd = "SELECT * FROM person where age = ?"
	// row := DBConnection.QueryRow(cmd, 20)
	// var p Person
	// err = row.Scan(&p.Name, &p.Age)
	// if err != nil {
	// 	if err == sql.ErrNoRows {
	// 		log.Println("Now row")
	// 	} else {
	// 		log.Println(err)
	// 	}
	// }
	// fmt.Println(p.Name, p.Age)

	// データの削除
	// cmd = "DELETE FROM person WHERE name = ?"
	// _, err = DBConnection.Exec(cmd, "Nancy")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
}
