package th_negroni

import (
	"github.com/urfave/negroni"
	"net/http"
	"fmt"
	"third-lib/th_negroni/middleware"
	"github.com/gorilla/mux"
  )
  
  func Run() {

	r := mux.NewRouter()
	r.HandleFunc("/host", HostHandler).Host("localhost")
  
	n := negroni.Classic()
	n.Use(middleware.NewMiddlerware1())
	n = n.With(middleware.NewMiddlerware2())

	api := mux.NewRouter()
	api.HandleFunc("/api/ver", ApiVerHandler)

	web := mux.NewRouter()
	web.HandleFunc("/web/ver", WebVerHandler)

	r.PathPrefix("/api").Handler(n.With(
		middleware.NewApiMiddlerware(),
		negroni.Wrap(api),
	))

	r.PathPrefix("/web").Handler(n.With(
		middleware.NewWebMiddlerware(),
		negroni.Wrap(web),
	))

	n.UseHandler(r)
	n.Run(":3030")
  }

  func HostHandler(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "the host is localhost")
}

func ApiVerHandler(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "api version: 1.0.0")
}

func WebVerHandler(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "web version: 1.0.0")
}