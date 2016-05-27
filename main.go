package main

import (
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
	"flag"
)

// struck to manage templates
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ =
			template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, nil)
}

func main() {

	var port = flag.String("port", ":8090", "The port addr of the app.")
	flag.Parse()

	http.Handle("/", &templateHandler{filename: "index.html"})
	http.Handle("/about", &templateHandler{filename: "about.html"})

	log.Println("Starting web server on ", *port)

	// start the web server
	if err := http.ListenAndServe(*port, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
