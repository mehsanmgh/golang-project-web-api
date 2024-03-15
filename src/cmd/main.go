package main

import (
	"log"

	"github.com/mehsanmgh/golang-project-web-api/api"
	"github.com/mehsanmgh/golang-project-web-api/config"
	"github.com/mehsanmgh/golang-project-web-api/data/cache"
	"github.com/mehsanmgh/golang-project-web-api/data/db"
)

func main() {
	cfg := config.GetConfig()
	
	err := cache.InitRedis(cfg)
	defer cache.CloseRedis()
	if err != nil {
		log.Fatal(err)
	}

	err = db.InitDb(cfg)
	defer db.CloseDb()
	if err != nil {
		log.Fatal(err)
	}
	

	api.InitServer(cfg)
}