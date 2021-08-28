package main

import (
	"log"
	"net/http"
	"text/template"
)

type Temps struct {
	notemp *template.Template
	index *template.Template
	hello *template.Template
}

func notemp() *template.Template {
	src := "<html><body><h1>NO TEMPLATE.</h1></body></html>"
	tmp, _ := template.New("index").Parse(src)
	return tmp
}

func setupTemp() *Temps {
	temps := new(Temps)

	temps.notemp = notemp()

	index, er := template.ParseFiles("templates/index.html")
	if er != nil {
		index = temps.notemp
	}
	temps.index = index

	hello, er := template.ParseFiles("templates/hello.html")
	if er != nil {
		hello = temps.notemp
	}
	temps.hello = hello

	return temps
}

func index(w http.ResponseWriter, r *http.Request, tmp *template.Template) {
	er := tmp.Execute(w, nil)
	if er != nil {
		log.Fatal(er)
	}
}

func hello(w http.ResponseWriter, r *http.Request, tmp *template.Template) {
	item := struct {
		Title string
		Message string
	} {
		Title: "Send values",
		Message: "This is Sample message.<br>これはサンプルです。",
	}

	er := tmp.Execute(w, item)

	if er != nil {
		log.Fatal(er)
	}
}

func main() {
	temps := setupTemp()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		index(w, r, temps.index)
	})
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		hello(w, r, temps.hello)
	})
	http.ListenAndServe("localhost:8080", nil)
}