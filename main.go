package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	cache "github.com/victorspringer/http-cache"
	"github.com/victorspringer/http-cache/adapter/memory"
)

func assetHandler(w http.ResponseWriter, r *http.Request) {
	assetPath := strings.TrimPrefix(r.URL.Path, "/asset/")
	if strings.HasSuffix(assetPath, ".png") || strings.HasSuffix(assetPath, ".jpeg") || strings.HasSuffix(assetPath, ".jpg") {
		webpPath := filepath.Join("./asset", strings.TrimSuffix(assetPath, filepath.Ext(assetPath))+".webp")
		if _, err := os.Stat(webpPath); err == nil {
			// log.Printf("%s -> %s", assetPath, webpPath)
			http.ServeFile(w, r, webpPath)
			return
		}
	}
	http.ServeFile(w, r, filepath.Join("./asset", assetPath))
}

func main() {
	memcached, err := memory.NewAdapter(
		memory.AdapterWithAlgorithm(memory.LRU),
		memory.AdapterWithCapacity(10000000),
	)
	if err != nil {
		log.Fatal(err)
	}

	cacheClient, err := cache.NewClient(
		cache.ClientWithAdapter(memcached),
		cache.ClientWithTTL(10*time.Minute),
		cache.ClientWithRefreshKey("opn"),
	)
	if err != nil {
		log.Fatal(err)
	}

	handler := cacheClient.Middleware(http.HandlerFunc(assetHandler))

	http.Handle("/asset/", handler)
	fmt.Println("Listening on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
