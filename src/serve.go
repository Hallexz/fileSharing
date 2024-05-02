package src

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func ServeFiles(w http.ResponseWriter, r *http.Request) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		http.Error(w, "Не удалось получить домашнюю директорию", http.StatusInternalServerError)
		return
	}
	path := r.URL.Path[len("/files/"):] // Используйте срез строки здесь
	fullPath := filepath.Join(homeDir, path)
	info, err := os.Stat(fullPath)
	if os.IsNotExist(err) {
		http.NotFound(w, r)
		return
	}
	if info.IsDir() {
		files, err := os.ReadDir(fullPath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		for _, file := range files {
			if file.IsDir() {

				fmt.Fprintf(w, "%s/\n", file.Name())
			} else {
				fmt.Fprintf(w, "%s\n", file.Name())
			}
		}
	} else {
		http.ServeFile(w, r, fullPath)
	}
}
