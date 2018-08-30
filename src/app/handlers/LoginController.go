package handlers

import (
	"app/models"
	"app/services"
	"encoding/json"
	"html/template"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	service := services.LoginService{}
	data := struct{ Error string }{""}

	if r.FormValue("register") != "" {
		http.Redirect(w, r, "/register-ui", 301)
	} else if r.FormValue("sign-in") != "" {
		err := service.Authenticate(r.FormValue("user-id"), r.FormValue("password"))
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
		http.Redirect(w, r, "/", 301)

	}

}

func LoginUIHandler(w http.ResponseWriter, r *http.Request) {
	responseTemplate := template.Must(template.ParseFiles("templates/login/login.html"))

	if err := responseTemplate.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func RegisterUIHandler(w http.ResponseWriter, r *http.Request) {
	responseTemplate := template.Must(template.ParseFiles("templates/register/register.html"))

	if err := responseTemplate.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func UsernameAvailability(w http.ResponseWriter, r *http.Request) {
	service := services.LoginService{}

	result, err := service.CheckUserAvailability(r.URL.Query().Get("username"))

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	encoder := json.NewEncoder(w)

	if err := encoder.Encode(result); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var userui models.UserUi
	err := decoder.Decode(&userui)
	if err == nil {
		user := models.User{Username: userui.Username}
		user.SetSecret(userui.Password)
		service := services.LoginService{}
		err := service.Register(user)
		if err != nil {
			http.Error(w, "Error registering the user", http.StatusInternalServerError)
		}
		encoder := json.NewEncoder(w)

		if err := encoder.Encode(true); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

	} else {
		http.Error(w, "Invalid data format", http.StatusInternalServerError)
	}
}
