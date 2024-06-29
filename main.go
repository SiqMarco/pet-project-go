package main

import (
	"fmt"
	"log"
	"net/http"
	"petPorject/api/src/config"
	"petPorject/api/src/routers"
)

func main() {
	config.Carregar()
	fmt.Printf("escutando na porta %d\n", config.Port)
	r := routers.GetRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
