package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"./db"
	"./routes"
	"./utils"
)

func main() {
	utils.ReadSettings()

	db.StartDbConnection()

	// listening for server's interrupt
	// not using it for some functionality for now,
	// can useful for solving some problems of server shutdown
	signals := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-signals

		log.Println(sig)

		done <- true
	}()

	go func() {
		<-done

		log.Println("Exit command received. Exiting ....")

		os.Exit(0)
	}()

	log.Println("Server started. Press Ctrl-C to stop server")

	//starting server
	routes.RunAllRoutes()
}
