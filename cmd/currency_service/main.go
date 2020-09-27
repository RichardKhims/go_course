package main

import (
	"fmt"
	"sync"

	"github.com/RichardKhims/go_course/internal/currency_service/config"
	"github.com/RichardKhims/go_course/internal/currency_service/database"
	"github.com/RichardKhims/go_course/internal/currency_service/server"
)

func main() {
	cfg, err := config.ReadConfig("configs/config.toml")
	if err != nil {
		panic(err)
	}
	fmt.Println(cfg)

	db, err := database.New(cfg.Database)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	wg := sync.WaitGroup{}
	wg.Add(1)
	updaterService := server.UpdaterService{
		Config: cfg.FcsApi,
		DB:     db,
	}
	go updaterService.Run(&wg)

	s := server.New(cfg.Server, db)
	fmt.Println(db)
	s.Run()
	wg.Wait()
}
