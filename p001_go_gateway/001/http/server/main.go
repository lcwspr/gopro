package main

import (
	"log"
	"net/http"
	"time"
)

const (
	Addr = ":8080"
)

func sayBye(w http.ResponseWriter, r *http.Request) {

	time.Sleep(1 * time.Second)
	w.Write([]byte("bye bye, this is httpServer"))
}

func main() {

	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime)

	mux := http.NewServeMux()

	mux.HandleFunc("/bye", sayBye)

	server := &http.Server{
		Addr:         Addr,
		WriteTimeout: time.Second * 3,
		Handler:      mux,
	}

	log.Println("Starting http srver at", Addr)
	log.Fatal(server.ListenAndServe())

}
