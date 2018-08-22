package handlers

import (
	"html/template"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	responseTemplate := template.Must(template.ParseFiles("templates/login/login.html"))

	r.ParseForm()

	if r.Form.Get("user-id") == "abc@abc" && r.Form.Get("password") == "test" {
		responseTemplate := template.Must(template.ParseFiles("templates/main/index.html"))

		if err := responseTemplate.ExecuteTemplate(w, "index.html", nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	data := struct{ Error string }{"Invalid Login..."}

	if err := responseTemplate.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
