package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
)

type Employee struct {
	Id       int
	Username string
	Password string
	Name     string
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "golang"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

var tmpl = template.Must(template.ParseGlob("form/*"))

func Index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM user")
	if err != nil {
		panic(err.Error())
	}
	emp := Employee{}
	res := []Employee{}
	for selDB.Next() {
		var id int
		var username, password, name string
		err = selDB.Scan(&id, &username, &password, &name)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.Name = name
		emp.Password = password
		emp.Username = username
		res = append(res, emp)
	}
	tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()
}

func usersJSON(response http.ResponseWriter, request *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM user")
	if err != nil {
		panic(err.Error())
	}
	emp := Employee{}
	res := []Employee{}
	for selDB.Next() {
		var id int
		var username, password, name string
		err = selDB.Scan(&id, &username, &password, &name)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.Name = name
		emp.Password = password
		emp.Username = username
		res = append(res, emp)
	}

	response.Header().Set("Content-type", "application/json")
	if request.Method == "GET" {
		var result, err = json.Marshal(res)

		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}
		response.Write(result)
		return
	}
	http.Error(response, "Page request failed", http.StatusBadRequest)
	defer db.Close()
}

func Show(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM user WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	emp := Employee{}
	for selDB.Next() {
		var id int
		var username, password, name string
		selDB.Scan(&id, &username, &password, &name)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.Name = name
		emp.Password = password
		emp.Username = username
	}
	tmpl.ExecuteTemplate(w, "Show", emp)
	defer db.Close()
}

func showByID(response http.ResponseWriter, request *http.Request){
	db := dbConn()
	nId := request.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM user WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	emp := Employee{}
	for selDB.Next() {
		var id int
		var username, password, name string
		selDB.Scan(&id, &username, &password, &name)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.Name = name
		emp.Password = password
		emp.Username = username
	}

	response.Header().Set("Content-type", "application/json")

	if request.Method == "GET" {
		result, err := json.Marshal(emp)
		if err != nil{
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}
		response.Write(result)
		return
	}

	http.Error(response, "ID not Found", http.StatusInternalServerError)
	defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM user WHERE id=?", nId)
	if err != nil {
		panic(err.Error())
	}
	emp := Employee{}
	for selDB.Next() {
		var id int
		var username, password, name string
		err = selDB.Scan(&id, &username, &password, &name)
		if err != nil {
			panic(err.Error())
		}
		emp.Id = id
		emp.Name = name
		emp.Password = password
		emp.Username = username
	}
	tmpl.ExecuteTemplate(w, "Edit", emp)
	defer db.Close()
}

func Insert(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		username := r.FormValue("username")
		password := r.FormValue("password")
		insForm, err := db.Prepare("INSERT INTO user(username, password, name) VALUES(?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(username, password, name)
		log.Println("INSERT: Name: " + name + " | Password: " + password + " | Username: " + username)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		username := r.FormValue("username")
		password := r.FormValue("password")
		id := r.FormValue("uid")
		insForm, err := db.Prepare("UPDATE user SET username=?, password=?, name=? WHERE id=?")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(username, password, name, id)
		log.Println("UPDATE: Name: " + name + " | Password: " + password + " | Username: " + username)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	emp := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM user WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(emp)
	log.Println("DELETE")
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func main() {
	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/", Index)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/new", New)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)
	http.HandleFunc("/jsonAll", usersJSON)
	http.HandleFunc("/show/id?", showByID)
	http.ListenAndServe(":8080", nil)
}
