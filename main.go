package main

import (
	"context"
	"fileSharing/src"
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Secret(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Secret information!\n")
}

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}

	db := client.Database("mydb")
	auth := src.NewAuth("myusername", "mypassword", db)

	http.HandleFunc("/", auth.BasicAuth(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, world!"))
	}))

	authenticator := src.NewAuth("admin", "admin", db)

	key := []byte("new-16-byte-key!!!!!!!!!")
	data := []byte("your-data-here")

	encryptedData, err := src.Encrypt(key, data)
	if err != nil {
		fmt.Println("Encryption error", err)
		return
	}
	fmt.Printf("Encrypted data: %x\n", encryptedData)

	http.HandleFunc("/secret", authenticator.BasicAuth(Secret))
	http.HandleFunc("/upload", src.UploadFile)
	http.HandleFunc("/files/", src.ServeFiles)
	http.ListenAndServe(":8080", nil)
}
