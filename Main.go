package main 


import (
	"fmt"
	"log"
	"net/http"
	"github.com/Akhenaten-Sama/go-gres/router"
)

func main(){
	route:= router.Router()
	log.Fatal(http.ListenAndServe(":8080", route))
	fmt.Println("starting route at port 8080")
}