package main

import (
	"io"
	"log"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte("<h1>Welcome to cloud native</h1>"))

	for k,v := range r.Header{
		for _, vv := range v{
			w.Header().Set(k,vv)
		}
	}
}

func main() {
	if err := http.ListenAndServe(":8080", mxu	); err != nil {
		log.Fatal("start server failed ,%s")
	}
}
