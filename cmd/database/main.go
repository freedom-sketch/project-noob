// This file is temporary and will be deleted along with the cmd/database folder
package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/freedom-sketch/sub2go/config"
	"github.com/freedom-sketch/sub2go/internal/database"
	"github.com/freedom-sketch/sub2go/internal/logger"
)

func main() {
	cfg, err := config.Load("config.json")
	if err != nil {
		log.Fatal("Failed to load config", err)
	}

	err = logger.Init(cfg.Logging)
	if err != nil {
		log.Fatal("Failed to initialize logging", err)
	}
	log := logger.Log

	if err := database.Connect(); err != nil {
		log.Fatal("Connect failed:", err)
	}
	log.Info("✅ Successful connection to the database")

	prettyJSON, err := json.MarshalIndent(cfg, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(prettyJSON))
}
