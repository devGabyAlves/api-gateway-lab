package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/users/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("users service ok"))
	})

	log.Println("[USERS] listening on :3001")
	http.ListenAndServe(":3001", nil)
}
