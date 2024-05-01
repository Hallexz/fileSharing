package main

import (
	"fileSharing/src"
	"net/http"
)

func main() {
	http.HandleFunc("/upload", src.UploadFile)
}
