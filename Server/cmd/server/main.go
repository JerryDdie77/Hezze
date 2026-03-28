package main

import (
	"log"

	"github.com/JerryDdie77/Hezze/Server/internal/config"
	"github.com/JerryDdie77/Hezze/Server/internal/database"
	"github.com/JerryDdie77/Hezze/Server/internal/handler"
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

	h := handler.NewHandler()
	r := handler.SetupRouter(h)
	
	if err := r.Run("8000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	

}
