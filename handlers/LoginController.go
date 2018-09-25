package handlers

import (
	"PlayingWithGoAndVue/helper"
	"PlayingWithGoAndVue/models"
	"PlayingWithGoAndVue/services"
	"encoding/json"
	"google.golang.org/appengine"
	"html/template"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var userUI models.UserUi
	err := decoder.Decode(&userUI)

	if err != nil {
		http.Error(w, "Invalid data", http.StatusInternalServerError)
	} else {
		service := services.LoginService{appengine.NewContext(r)}
		err := service.Authenticate(userUI.Username, userUI.Password)
		if err != nil {
			http.Error(w, "Invalid user/password", http.StatusInternalServerError)
		} else {
			loginHelper := helper.LoginHelper{}
			err := loginHelper.SetSessionValues(w, r, userUI)
			if err != nil {
				http.Error(w, "Error generating Session", http.StatusInternalServerError)
			}
		}
	}
}

func LoginUIHandler(w http.ResponseWriter, r *http.Request) {
	loginHelper := helper.LoginHelper{}

	if ok := loginHelper.RedirectIfLoggedIn(w, r); ok {
		return
	}

	responseTemplate := template.Must(template.ParseFiles("./static/login/login.html"))

	if err := responseTemplate.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func RegisterUIHandler(w http.ResponseWriter, r *http.Request) {
	responseTemplate := template.Must(template.ParseFiles("./static/register/register.html"))

	if err := responseTemplate.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func UsernameAvailability(w http.ResponseWriter, r *http.Request) {
	service := services.LoginService{appengine.NewContext(r)}

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
	var userUI models.UserUi
	err := decoder.Decode(&userUI)
	if err == nil {
		user := models.User{Username: userUI.Username}
		user.SetSecret(userUI.Password)
		service := services.LoginService{appengine.NewContext(r)}
		err := service.Register(&user)
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

func VerifyUser(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	loginHelper := helper.LoginHelper{}
	if loginHelper.IsExempt(r.URL.Path) {
		next(w, r)
		return
	}

	if username := loginHelper.GetSessionValues(r); username != "" {
		loginService := services.LoginService{appengine.NewContext(r)}
		if ok, _ := loginService.CheckUserAvailability(username); ok {
			next(w, r)
			return
		}
	}

	http.Redirect(w, r, "/login-ui", http.StatusTemporaryRedirect)
}