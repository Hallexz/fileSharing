package src

import (
	"io"
	"net/http"
	"os"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	// Получаем файл из запроса
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Не удалось получить файл", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Создаем новый файл на сервере
	f, err := os.OpenFile("./uploads/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Не удалось создать файл", http.StatusInternalServerError)
		return
	}
	defer f.Close()

	// Копируем содержимое файла
	if _, err := io.Copy(f, file); err != nil {
		http.Error(w, "Не удалось скопировать содержимое файла", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Файл успешно загружен"))
}
