package main

import(
    "log"
     "net/http"
     "text/template"
     "path/filepath"
     "sync"
    )
    
// struck to manage templates
type templateHandler struct {
   once  sync.Once
   filename string
   templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
        t.once.Do(func() {
              t.templ = 
              template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
        })
        t.templ.Execute(w,nil)
}
    
func main() {

http.Handle("/", &templateHandler{ filename: "index.html"})
http.Handle("/about", &templateHandler{ filename: "about.html"})

// start the web server
    if err := http.ListenAndServe(":80", nil); err != nil {
            log.Fatal("ListenAndServe:", err)
    }
}