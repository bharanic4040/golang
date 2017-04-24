package main

import(
"net/http"
"fmt"
)

func Logger(inner http.Handler,name string) http.Handler{

return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
fmt.Println("Inside decorator")

inner.ServeHTTP(w,r)


})

}
