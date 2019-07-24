package app

import (
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

var err error

func Run() {

	router := httprouter.New()

	router.GET("/", Index)
	router.POST("/compile", Compile)

	router.ServeFiles("/assets/*filepath", http.Dir("assets"))

	log.Println("Server listening at port 8100")
	log.Fatal(http.ListenAndServe(":8100", router))

}
