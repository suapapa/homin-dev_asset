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
			path = path[1:]
		} else {
			path = "."
		}

		path = strings.TrimPrefix(path, "/asset/")
		if strings.HasSuffix(path, ".png") || strings.HasSuffix(path, ".jpeg") || strings.HasSuffix(path, ".jpg") {
			tempPath := filepath.Join("./asset", strings.TrimSuffix(path, filepath.Ext(path))+".webp")
			if _, err := os.Stat(tempPath); err == nil {
				// log.Printf("%s -> %s", path, webpPath)
				// http.ServeFile(w, r, webpPath)
				// return
				path = tempPath
			}
		}

		err := cache.WriteFile(w, path)
		if err == nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		} else if err == filecache.ItemIsDirectory {
			http.ServeFile(w, r, path)
		}
	}

	// memcached, err := memory.NewAdapter(
	// 	memory.AdapterWithAlgorithm(memory.LRU),
	// 	memory.AdapterWithCapacity(10_000_000),
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// cacheClient, err := cache.NewClient(
	// 	cache.ClientWithAdapter(memcached),
	// 	cache.ClientWithTTL(10*time.Minute),
	// 	cache.ClientWithRefreshKey("opn"),
	// )
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// handler := cacheClient.Middleware(http.HandlerFunc(assetHandler))

	http.Handle("/asset/", http.HandlerFunc(fileServer))
	fmt.Println("Listening on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
