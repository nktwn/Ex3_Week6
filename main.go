package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type PageData struct {
	Title        template.HTML
	BusinessName string
	Slogan       string
	Timestamp    string
}

func index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		data := PageData{
			Title:        "Ex3_week3 | AI & GPT",
			BusinessName: "Business,",
			Slogan:       "we get things done!",
			Timestamp:    time.Now().Format(time.RFC1123),
		}
		t, err := template.ParseFiles("template.html")
		if err != nil {
			http.Error(w, fmt.Sprintf("Error parsing template: %v", err), http.StatusInternalServerError)
			return
		}
		err = t.Execute(w, data)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error executing template: %v", err), http.StatusInternalServerError)
			return
		}
	})
}

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		next.ServeHTTP(w, r)
		fmt.Printf("Request processed in %v\n", time.Since(startTime))
	})
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", logger(index()))
	fmt.Println("Server is running at http://localhost:8080/")
	http.ListenAndServe(":8080", mux)
}
