package main

import (
	"net/http"
	"fmt"
	"html/template"
)

func main(){
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var data = map[string]string{
			"Name":    "FAOZI",
			"Message": "have a nice day",
		}

		var t, err = template.ParseFiles("resources/views/template.html")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		t.Execute(w, data)
	})

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
