package configuration

import "net/http"

type Security struct {
}

func (s *Security) VerifyUser(w http.ResponseWriter, r *http.Request) {

}
