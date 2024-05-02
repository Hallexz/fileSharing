package src

import (
	"context"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Auth struct {
	username string
	password string
	db       *mongo.Database
}

func NewAuth(username, password string, db *mongo.Database) *Auth {
	return &Auth{username: username, password: password, db: db}
}

func (a *Auth) BasicAuth(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()

		if !ok || !a.validateCredentials(user, pass) {
			w.Header().Set("WWW-Authenticate", `Basic realm="Enter login and password"`)
			w.WriteHeader(401)
			w.Write([]byte("You are not authorized to view this page\n"))
			return
		}
		handler(w, r)
	}
}

func (a *Auth) validateCredentials(username, password string) bool {
	var result bson.M
	err := a.db.Collection("users").FindOne(context.Background(), bson.M{"username": username}).Decode(&result)
	if err != nil {
		return false
	}

	return result["password"] == password
}
