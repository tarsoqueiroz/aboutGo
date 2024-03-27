package main

import (
	"fmt"
	"net/http"
)

func handlRoot(w http.ResponseWriter, r *http.Request) {
	strMessage := "Hello World from Go!!!"

	fmt.Println(strMessage)
	fmt.Fprint(w, strMessage)
}

func handlTest(w http.ResponseWriter, r *http.Request) {
	strMessage := "Testing Hello World..."

	fmt.Println(strMessage)
	fmt.Fprint(w, strMessage)
}

func main() {

	http.HandleFunc("/", handlRoot)
	http.HandleFunc("/test", handlTest)

	fmt.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)

}
