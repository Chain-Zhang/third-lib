package middleware

import(
	"fmt"
	"net/http"
)
type ApiMiddlerware struct{

}

func NewApiMiddlerware() *ApiMiddlerware {
	return &ApiMiddlerware{}
}

func (self *ApiMiddlerware)ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc){
	fmt.Println("arrive at api middleware")
	next(rw, r)
}