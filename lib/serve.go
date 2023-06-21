package lib

import (
	"log"
	"net/http"
)

func Serve(file string, port string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, file)
	})
	log.Println("Serving resume on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
