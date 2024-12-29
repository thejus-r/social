package main

import (
	"log"

	"github.com/thejus-r/social/package/db"
	"github.com/thejus-r/social/package/store"
)

func main() {
	addr := "postgres://admin:adminpassword@localhost/social?sslmode=disable"
	conn, err := db.New(addr, 3, 3, "15m")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	store := store.NewStorage(conn)
	db.Seed(store)
}
