package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request received")
	defer log.Println("Request processing completed")

	select {
	case <-time.After(5 * time.Second):
		log.Println("Processing completed successfully")
		w.Write([]byte("Request processed successfully"))
	case <-ctx.Done():
		log.Println("Request cancelled by client")
		http.Error(w, "Request cancelled", http.StatusRequestTimeout)
	}
}
