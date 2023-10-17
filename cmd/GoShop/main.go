package main

import (
	"GoShop/pkg/db"
	"log"
	"net/http"
)

func main() {
	err := db.InitDb("root:root@tcp(127.0.0.1:3307)/GoShop")

	if err != nil {
		log.Fatal("Connection error!", err)
		return
	}

	http.ListenAndServe("localhost:8082", nil)
}
