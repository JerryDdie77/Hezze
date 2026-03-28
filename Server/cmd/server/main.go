package main

import (
	"log"

	"github.com/JerryDdie77/Hezze/Server/internal/config"
	"github.com/JerryDdie77/Hezze/Server/internal/database"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("GetConfig: %v", err)
	}

	dsn := cfg.DSN()
	
	db, err := database.NewPostgresDB(dsn)
	if err != nil {
		log.Fatalf("NewPostgreDB: %v", err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		log.Fatalf("Ping: %v", err)
	}

	

}
