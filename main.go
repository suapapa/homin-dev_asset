package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func assetHandler(w http.ResponseWriter, r *http.Request) {
	assetPath := strings.TrimPrefix(r.URL.Path, "/asset/")
	if strings.HasSuffix(assetPath, ".png") || strings.HasSuffix(assetPath, ".jpeg") || strings.HasSuffix(assetPath, ".jpg") {
		webpPath := filepath.Join("./asset", strings.TrimSuffix(assetPath, filepath.Ext(assetPath))+".webp")
		if _, err := os.Stat(webpPath); err == nil {
			log.Printf("%s -> %s", assetPath, webpPath)
			http.ServeFile(w, r, webpPath)
			return
		}
	}
	http.ServeFile(w, r, filepath.Join("./asset", assetPath))
}

func main() {
	http.HandleFunc("/asset/", assetHandler)
	http.ListenAndServe(":8080", nil)
}
