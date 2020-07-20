package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main()  {
	log.Printf("My Server Started")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":6767", nil))
}

func handler(w http.ResponseWriter, r *http.Request)  {
	ctx := r.Context()
	log.Printf("Handler started")
	defer log.Printf("handler ended")
	// making the server to slow down, so that we can see the results
	select {
	case <- time.After(5 * time.Second):
		fmt.Fprintf(w, "You are running a server")
	case <- ctx.Done():
		err := ctx.Err()
		log.Print(ctx.Err())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}