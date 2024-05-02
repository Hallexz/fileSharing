package src

import (
	"net/http"
)

type Auth struct {
	username string
	password string
}

func NewAuth(username, password string) *Auth {
	return &Auth{username: username, password: password}
}

func (a *Auth) BasicAuth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()

		if !ok || user != a.username || pass != a.password {
			w.Header().Set("WWW-Authenticate", `Basic realm="Введите логин и пароль"`)
			w.WriteHeader(401)
			w.Write([]byte("Вы не авторизованы для просмотра этой страницы.\n"))
			return
		}

		handler(w, r)
	}
}
