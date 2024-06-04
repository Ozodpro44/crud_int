package main

import (
	"app/config"
	"app/pkg/db"
	"fmt"
	"log"
)

func main() {

	log.Println("starting proccess...")

	cfg := config.NewConfig()
	log.Println("successfully loaded config!")

	con, err := db.Conn(cfg)
	if err != nil {
		log.Println("error in connecting with database:", err)
	}
	log.Println("successfully connected with database!")

	fmt.Println(con)

}
