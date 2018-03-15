package middleware

import(
	"fmt"
	"net/http"
)

type WebMiddlerware struct{

}

func NewWebMiddlerware() *WebMiddlerware {
	return &WebMiddlerware{}
}

func (self *WebMiddlerware)ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc){
	fmt.Println("arrive at web middleware")
	next(rw, r)
}