package main

import "net/http"

// /asset 의 경로의 파일들을 파일서빙하는 서버

func main() {
	// /asset 경로의 파일들을 서빙하는 서버를 만듭니다.
	// http.FileServer를 사용하여 파일을 서빙합니다.
	// http.StripPrefix를 사용하여 /asset 경로를 제거합니다.
	// http.ListenAndServe를 사용하여 서버를 실행합니다.

	http.Handle("/asset/", http.StripPrefix("/asset/", http.FileServer(http.Dir("./asset"))))
	http.ListenAndServe(":8080", nil)
}
