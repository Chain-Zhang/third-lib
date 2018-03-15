package middleware

import(
	"fmt"
	"net/http"
)

type Middlerware1 struct{

}

func NewMiddlerware1() *Middlerware1 {
	return &Middlerware1{}
}

func (self *Middlerware1)ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc){
	fmt.Println("arrive at middleware 1")
	next(rw, r)
}