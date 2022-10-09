package main

import (
	//"database/sql"
	"fmt"
	"log"
	"net/http"

	//"github.com/gorilla/mux"
	//_ "github.com/go-sql-driver/mysql"
	"webapp/src/config"
	"webapp/src/router"
)

func main() {
	config.Carregar()
	fmt.Println("API rodando")

	r := router.Gerar()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
