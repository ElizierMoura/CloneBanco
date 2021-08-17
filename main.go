package main

import (
	"apiGetBanco/controllers/controled"
	"apiGetBanco/src/router"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Iniciando api")
	router := router.GerarRouter()

	router.HandleFunc("/clonar", controled.GetBanco)

	http.ListenAndServe(":8080", router)
}
