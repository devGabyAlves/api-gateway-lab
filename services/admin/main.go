package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/admin/secret", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("FLAG{internal_admin_secret}"))
	})

	log.Println("[ADMIN] listening on :3002")
	http.ListenAndServe(":3002", nil)
}
