package main

import (
	"errors"
	"log"
)

func loadConfig() error {
	return errors.New("config file not found")
}

func connect() error {
	return errors.New("database not reachable")
}

func loadUser() error {
	return nil
}

func main() {
	log.Println("Starting service")
	err := loadUser()
	if err != nil {
		log.Fatal("Critical failure, shutting down:", err)
	}
	log.Println("Service started successfully")
}
