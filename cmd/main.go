package main

import (
	"fmt"

	"github.com/jasurxaydarov/marifat_ac_backend/api"
	"github.com/jasurxaydarov/marifat_ac_backend/config"
	"github.com/jasurxaydarov/marifat_ac_backend/pgx/db"
	"github.com/jasurxaydarov/marifat_ac_backend/storage"
)

func main() {

	fmt.Println("start")

	cfg := config.Load()

	fmt.Println(cfg)

	pgxCoon, err := db.ConnectDB(cfg.PgConfig)

	if err != nil {
		fmt.Println("err on conn db", err)
		return
	}
	fmt.Println(pgxCoon)

	str := storage.NewStorage(pgxCoon)

	api.Api(str)
}
