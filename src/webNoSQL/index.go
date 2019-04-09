package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
)

type student struct {
	Nama  string `bson:"name"`
	Grade int    `bson:"Grade"`
}

func connect() (*mgo.Session, error) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		fmt.Println(session)
		fmt.Println("kesini")
		fmt.Println(err.Error())
		return nil, err
	}
	fmt.Println("disini")
	return session, nil
}

func insert() {
	var session, err = connect()
	if err != nil {
		fmt.Println("Error!", err.Error()+" error disini")
		return
	}
	defer session.Close()

	var collection = session.DB("belajar_golang").C("student")
	err = collection.Insert(&student{"Wick", 2}, &student{"Ethan", 2})
	if err != nil {
		fmt.Println("Error!", err.Error()+" disini deng")
		return
	}

	fmt.Println("Insert success!")
}

func main() {
	insert()
}
