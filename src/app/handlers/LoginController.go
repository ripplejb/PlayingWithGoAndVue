package handlers

import (
	"app/models"
	"app/services"
	"html/template"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	user := models.User{Username: r.FormValue("user-id")}
	user.SetSecret(r.FormValue("password"))
	service := services.LoginService{}
	data := struct{ Error string }{""}

	if r.FormValue("register") != "" {
		err := service.Register(user)
		if err != nil {
			data.Error = "Error while registering the user."
		}
	} else if r.FormValue("sign-in") != "" {
		err := service.Authenticate(user)
		if err != nil {
			data.Error = "Invalid Login."
		}
	}

	if data.Error != "" {
		responseTemplate := template.Must(template.ParseFiles("templates/login/login.html"))

		if err := responseTemplate.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	} else {
		responseTemplate := template.Must(template.ParseFiles("templates/main/index.html"))

		if err := responseTemplate.Execute(w, nil); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	}

}
