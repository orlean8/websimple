package main

import (
	"html/template"
	"net/http"
	"fmt"
	"../database"

)

type Page struct {
	Name string
	Val int
	Status string
}

func main (){

	temples := template.Must(template.ParseFiles("templates/index.html"))
	inMemDB := database.CreateNewDB()

	http.HandleFunc( "/get", func(w http.ResponseWriter, r *http.Request) {

		p:=Page {Name: "default"}

		if name:=r.FormValue("name"); name != "" {
			p.Name=name
			p.Val, p.Status = inMemDB.Get(name)
		}

		if err:= temples.ExecuteTemplate(w, "index.html", p); err!=nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	} )

	http.HandleFunc( "/set", func(w http.ResponseWriter, r *http.Request) {


		if name:=r.FormValue("name"); name != "" {
			inMemDB.Set(name)
		}

		fmt.Errorf("added")

	} )


	fmt.Println(http.ListenAndServe(":8080",nil))
}