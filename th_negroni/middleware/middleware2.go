package middleware

import(
	"fmt"
	"net/http"
	_ "github.com/urfave/negroni"
)

type Middlerware2 struct{

}

func NewMiddlerware2() *Middlerware2 {
	return &Middlerware2{}
}

func (self *Middlerware2)ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc){
	fmt.Println("arrive at middleware 2")
	next(rw, r)
}