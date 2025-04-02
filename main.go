package main

import (
	"log"

	"github.com/EmmanoelDan/agro-pop/config"
	"github.com/EmmanoelDan/agro-pop/routers"
	_ "github.com/lib/pq"
)

func main() {
	_, err := config.ConnectDatabase()
	if err != nil {
		log.Fatal(err)
	}

	routers.Init()
}
