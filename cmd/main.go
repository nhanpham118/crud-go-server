package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
)

func init()  {
	// Load .env file 
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}
}

func main() {
	log.Println("Starting server ...")

	// Gracefully exit on keyboard interupt
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	app, _ := initGoApp()
	go app.Start()

	<-c
	log.Print("Shutting down")
	os.Exit(0)
}