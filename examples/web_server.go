package examples

import (
	"log"
	"net/http"
)

func helloWeb(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Web!"))
}

func RunWebServer() {
	http.HandleFunc("/", helloWeb)
	err := http.ListenAndServe("localhost:9999", nil)
	if err != nil {
		log.Println("listen and serve:", err)
	}
}
