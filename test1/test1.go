package main

import (
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func home(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("html/home.html"))

	data := make(map[string]string)
	data["Aboutme"] = "About me"
	data["Description"] = "Personal Description"

	tmpl.Execute(w, data)
}

func greeting(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("html/greeting.html"))

	data := make(map[string]string)
	data["Time"] = time.Now().Format("11:29:07")
	data["Day"] = time.Now().Weekday().String()

	tmpl.Execute(w, data)
}

func random(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("html/random.html"))

	data := make(map[int]string)

	data[0] = "Failure is just success in progress"
	data[1] = "F.A.I.L. - First Attempt in Learning, don't give up"
	data[2] = "Success is on the same road as failure; success is just a little further down the road"
	data[3] = "Decide. Commit. Succeed."
	data[5] = "Accept failure, everyone fails at something but don't accept not trying"
	data[5] = "There are no secretes to success, it is the result of preperation, hard work and learning from failure"

	quote := data[rand.Intn(6)]

	tmpl.Execute(w, quote)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/greeting", greeting)
	mux.HandleFunc("/random", random)

	fs := http.FileServer(http.Dir("./css"))

	mux.Handle("/css/", http.StripPrefix("/css/", fs))

	log.Println("Starting server on Port 8000...")
	log.Println("See page at http://localhost:8000/")

	err := http.ListenAndServe(":8000", mux)

	if err != nil {
		panic(err)
	}
}
