package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main()  {
	r := chi.NewRouter()

	fmt.Println("API is running...")
	http.ListenAndServe(":3000", r)
}