package main

import(
	"fmt"
	"net/http"
	"encoding/json"
)

type Students struct{
	ID string
	Nama string
	NIM string
}

var students = []Students{
	Students{"1", "Mei Rusfandi", "123"},
	Students{"2", "Mei", "758"},
	Students{"3", "Rusfandi", "246"},
	Students{"4", "Rusfandi Mei", "043"},
	Students{"5", "Mei Mei", "234"},
}

func users(response http.ResponseWriter, request *http.Request){
	response.Header().Set("Content-Type", "application/json")

	if request.Method == "POST"{
		var result, err = json.Marshal(students)

		if err != nil{
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}
		response.Write(result)
		return
	}

	http.Error(response, "Page request failed", http.StatusBadRequest)
}

func user(response http.ResponseWriter, request *http.Request){
	response.Header().Set("Content-Type", "application/json")

	if request.Method == "POST"{
		var id = request.FormValue("id")
		var result []byte
		var err error

		for _, i := range students{
			if i.ID == id{
				result, err = json.Marshal(i)

				if err != nil{
					http.Error(response, err.Error(), http.StatusInternalServerError)
				}

				response.Write(result)
				return
			}
		}
	}

	http.Error(response, "user Not Found", http.StatusBadRequest)
	return
}

func main(){
	http.HandleFunc("/users", users)
	http.HandleFunc("/user", user)

	fmt.Println("Starting online on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}