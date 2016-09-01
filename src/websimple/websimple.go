package main

import (
	"html/template"
	"net/http"
	"fmt"
	"../database"

)


type Page struct {
	Records []database.ValuePair
	Status string
}

func main (){

	temples := template.Must(template.ParseFiles("templates/set.html", "templates/index.html", "templates/result.html"))
	inMemDB := database.CreateNewDB()
	p := Page{}

	http.HandleFunc( "/get", func(w http.ResponseWriter, r *http.Request) {


		if name:=r.FormValue("name"); name != "" {
			val, status := inMemDB.Get(name)
			p.Records=append([]database.ValuePair{},database.ValuePair{name,val})
			p.Status = status

		} else {
			if all :=r.FormValue("all"); all == "true" {
				p.Records, p.Status = inMemDB.GetAll()
							}
		}

		if err:= temples.ExecuteTemplate(w, "result.html", p); err!=nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	} )

	http.HandleFunc( "/set", func(w http.ResponseWriter, r *http.Request) {


		if name:=r.FormValue("name"); name != "" {
			inMemDB.Set(name)
		}

		if err:= temples.ExecuteTemplate(w, "set.html", nil); err!=nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}


	} )

	http.HandleFunc( "/start", func(w http.ResponseWriter, r *http.Request) {


		if name:=r.FormValue("name"); name != "" {
			inMemDB.Set(name)
		}

		if err:= temples.ExecuteTemplate(w, "index.html", nil); err!=nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}


	} )

	fmt.Println(http.ListenAndServe(":8080",nil))
}