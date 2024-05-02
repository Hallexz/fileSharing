package main

import (
	"fileSharing/src"
	"fmt"
	"net/http"
)

func Secret(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Secret information!\n")
}

func main() {
	authenticator := src.NewAuth("admin", "admin")
	http.HandleFunc("/secret", authenticator.BasicAuth(Secret))
	http.HandleFunc("/upload", src.UploadFile)
	http.HandleFunc("/files/", src.ServeFiles)
	http.ListenAndServe(":8080", nil)
}
