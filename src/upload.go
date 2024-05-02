package src

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"io"
	"net/http"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to receive file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		http.Error(w, "Failed to connect to MongoDB", http.StatusInternalServerError)
		return
	}
	defer client.Disconnect(context.Background())

	collection := client.Database("mydb").Collection("files")

	fileContent, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		return
	}

	doc := bson.M{
		"filename": handler.Filename,
		"content":  fileContent,
	}

	_, err = collection.InsertOne(context.Background(), doc)
	if err != nil {
		http.Error(w, "Failed to save file to MongoDB", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("File uploaded successfully"))
}
