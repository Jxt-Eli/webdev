package main

import(
	"log"
	"net/http"
)

func middleware(f http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		log.Println(r.URL.Path)
		f.ServeHTTP(w, r)
	})
}
