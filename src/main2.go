package main

import(
  "net/http"
  "text/template"
  "github.com/go-zoo/bone"
  "assets"
  "github.com/unrolled/render"
)

func main () {
  r := render.New(render.Options{
    Directory: "views",
    Extensions: []string{".html"},
  })
  ServeResource := assets.ServeResource
  mux := bone.New()
  mux.Get("/home/:id", http.HandlerFunc(homeController))
  mux.Handle("/foo", http.HandlerFunc(fooController))
  mux.Handle("/bar", http.HandlerFunc(barController))
	mux.Handle("/", http.HandlerFunc(rootController))
  mux.HandleFunc("/img/", ServeResource)
	mux.HandleFunc("/css/", ServeResource)
  mux.HandleFunc("/js/", ServeResource)

  func rootController(rw http.ResponseWriter, req *http.Request) {
    r.HTML(rw, 200, "index", nil)
  }

  func homeController(rw http.ResponseWriter, req *http.Request) {
    id := bone.GetValue(req, "id")
  	view, _ := template.ParseFiles("views/index.html", "views/_footer.html", "views/_header.html")
    view.Execute(rw, id)
  }

  func fooController(rw http.ResponseWriter, req *http.Request) {
  	view, _ := template.ParseFiles("views/foo.html", "views/_footer.html", "views/_header.html")
    view.Execute(rw, nil)
  }

  func barController(rw http.ResponseWriter, req *http.Request) {
  	view, _ := template.ParseFiles("views/bar.html", "views/_footer.html", "views/_header.html")
    view.Execute(rw, nil)
  }

  http.ListenAndServe(":8080", mux)
}
