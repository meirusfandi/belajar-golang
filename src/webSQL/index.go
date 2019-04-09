package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID       string
	USERNAME string
	PASSWORD string
	NAME     string
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/golang")

	if err != nil {
		return nil, err
	}

	return db, nil
}

func sqlQuery() {
	db, err := connect()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	rows, err := db.Query("select id, username, password, name from user")

	if err != nil {
		fmt.Println(err.Error() + " get data failed")
		return
	}
	defer rows.Close()

	var result []User

	for rows.Next() {

		var i = User{}

		var err = rows.Scan(&i.ID, &i.USERNAME, &i.PASSWORD, &i.NAME)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, i)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, i := range result {
		fmt.Println(i.NAME+" "+i.USERNAME+" "+i.PASSWORD)
	}
}

func sqlQueryRow() {
	var db, err = connect()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// db.Close()

	var result = User{}
	err = db.QueryRow("select username, password, name from user").Scan(&result.USERNAME, &result.PASSWORD, &result.NAME)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// fmt.Println("Username : %s",result.USERNAME, "\nPassword : %s",result.PASSWORD , "\nName : %s", result.NAME)
	fmt.Printf("Username : %s\nPassword : %s\nName : %s", result.USERNAME, result.PASSWORD, result.NAME)
}

func sqlPrepare() {
	db, err := connect()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// db.Close()

	statement, err := db.Prepare("select id, username, name from user where id = ?")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var result1 = User{}
	statement.QueryRow("1").Scan(&result1.ID, &result1.USERNAME, &result1.NAME)
	fmt.Printf("Username : %s\nName : %s\n", result1.USERNAME, result1.NAME)

	var result2 = User{}
	statement.QueryRow("2").Scan(&result2.ID, &result2.USERNAME, &result2.NAME)
	fmt.Printf("Username : %s\nName : %s\n", result2.USERNAME, result2.NAME)

	var result3 = User{}
	statement.QueryRow("3").Scan(&result3.ID, &result3.USERNAME, &result3.NAME)
	fmt.Printf("Username : %s\nName : %s\n", result3.USERNAME, result3.NAME)

	var result4 = User{}
	statement.QueryRow("4").Scan(&result4.ID, &result4.USERNAME, &result4.NAME)
	fmt.Printf("Username : %s\nName : %s\n", result4.USERNAME, result4.NAME)
}

func sqlCRUD() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//sql insert query
	_, err = db.Exec("insert into user values (?, ?, ?, ?)", "5", "coro", "coro", "coroooo")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("insert success!!")

	//sql update query
	_, err = db.Exec("update user set name = ? where id = ?", "cembre gatel", "5")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Update Success!!")

	//sql delete query
	_, err = db.Exec("delete from user where id = ?", "5")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Delete Success!")

}

func main() {
	sqlQuery()
	// sqlQueryRow()
	// sqlPrepare()
	// sqlCRUD()
}
