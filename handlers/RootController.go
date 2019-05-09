package handlers

import (
	"html/template"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {

	responseTemplate := template.Must(template.ParseFiles("./templates/main/index.html"))

	if err := responseTemplate.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
