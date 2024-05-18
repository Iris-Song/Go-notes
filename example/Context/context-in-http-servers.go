package main

import (
	"fmt"
	"net/http"
	"time"
)

// A Context carries deadlines, cancellation signals,
// and other request-scoped values across API boundaries and goroutines.

func hello(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")
	// While working, keep an eye on the context’s Done() channel for a signal that we should cancel the work and return as soon as possible.
	case <-ctx.Done():
		// The context’s Err() method returns an error that explains why the Done() channel was closed.
		err := ctx.Err()
		fmt.Println("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}
