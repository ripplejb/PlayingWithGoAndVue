package helper

import (
	"github.com/gorilla/sessions"
	"models"
	"net/http"
	"strings"
)

type LoginHelper struct {
}

var (
	key         = []byte("3f3db226-af7a-11e8-96f8-529269fb1459")
	store       = sessions.NewCookieStore(key)
	sessionName = "edd4d5c4-afc7-11e8-96f8-529269fb1459"
)

func (lh *LoginHelper) GetSessionValues(r *http.Request) string {
	var userName string

	if session, err := store.Get(r, sessionName); err == nil {

		if val := session.Values["user"]; val != nil {
			userName = val.(string)
		}
	}

	return userName
}

func (lh *LoginHelper) SetSessionValues(w http.ResponseWriter, r *http.Request, userUI models.UserUi) error {
	if session, err := store.Get(r, sessionName); err != nil {
		return err
	} else {
		session.Values["user"] = userUI.Username
		session.Save(r, w)
	}
	return nil
}

func (lh *LoginHelper) RedirectIfLoggedIn(w http.ResponseWriter, r *http.Request) bool {

	userName := lh.GetSessionValues(r)
	if userName != "" {
		http.Redirect(w, r, "/", http.StatusFound)
		return true
	}
	return false
}

func (lh *LoginHelper) IsExempt(url string) bool {
	urlList := [...]string{"/login", "/login-ui", "/register", "/register-ui", "/assets/"}

	for _, u := range urlList {
		if strings.HasPrefix(url, u) {
			return true
		}
	}
	return false
}
