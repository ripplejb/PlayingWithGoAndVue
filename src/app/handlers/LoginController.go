package handlers

import (
	"html/template"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	responseTemplate := template.Must(template.ParseFiles("templates/login/login.html"))

	data := struct{ Error string }{"Invalid Login..."}

	if err := responseTemplate.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
