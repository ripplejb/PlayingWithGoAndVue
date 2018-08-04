package handlers

import (
	"app/models"
	"html/template"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	p := models.Page{Name: "Go Ripal"}
	if name := r.FormValue("name"); name != "" {
		p.Name = name
	}

	responseTemplate := template.Must(template.ParseFiles("templates/main/index.html"))

	if err := responseTemplate.ExecuteTemplate(w, "index.html", p); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
