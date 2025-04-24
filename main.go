package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gokyle/filecache"
)

func main() {
	cache := filecache.NewDefaultCache()
	cache.MaxSize = 10 * filecache.Megabyte
	cache.Start()

	fileServer := func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if len(path) > 1 {
			path = path[1:] // trim leading slash
		} else {
			path = "."
		}

		path = strings.TrimPrefix(path, "asset/")
		path = filepath.Join("./asset", path)
		if strings.HasSuffix(path, ".png") || strings.HasSuffix(path, ".jpeg") || strings.HasSuffix(path, ".jpg") {
			webpPath := strings.TrimSuffix(path, filepath.Ext(path)) + ".webp"
			if _, err := os.Stat(webpPath); err == nil {
				if err := cache.WriteFile(w, webpPath); err != nil {
					log.Printf("Error writing file to cache, %s: %v", path, err)
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				} else {
					return
				}
			}
		}

		if err := cache.WriteFile(w, path); err != nil {
			log.Printf("Error writing file to cache, %s: %v", path, err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}

	http.Handle("/asset/", http.HandlerFunc(fileServer))
	fmt.Println("Listening on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
