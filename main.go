package main

import (
	"github.com/bmizerany/pat"
	"html/template"
	"log"
	"net/http"
)

func main() {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(home))
	mux.Post("/", http.HandlerFunc(send))
	mux.Get("/confirmation", http.HandlerFunc(confirmation))

	log.Println("Listening...")
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	render(w, "index.html", nil)
}

func send(w http.ResponseWriter, r *http.Request) {
	msg := &Message{
		email:    r.PostFormValue("email"),
		password: r.PostFormValue("password"),
	}

	if msg.Validate() == false {
		render(w, "index.html", msg)
		return
	}
	if err := msg.Deliver(); err != nil {
		log.Println(err)
		http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "confirmation.html", http.StatusSeeOther)
}

func confirmation(w http.ResponseWriter, r *http.Request) {
	render(w, "confirmation.html", nil)
}

func render(w http.ResponseWriter, filename string, data interface{}) {
	tmpl, err := template.ParseFiles(filename)
	if err != nil {
		log.Println(err)
		http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Println(err)
		http.Error(w, "Sorry, something went wrong", http.StatusInternalServerError)
	}
}
