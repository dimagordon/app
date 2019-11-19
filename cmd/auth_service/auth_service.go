package main

import (
	"app/server"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	app := server.NewApplication()

	log.Print("Starting Server.")

	go func() {
		if err := app.Run(); err != nil {
			log.Fatal(err)
		}
	}()

	waitSignal()
	log.Print("Gracefully shutting down app, press CTRL + C to force exit")
	app.Stop()
}

func waitSignal() {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGINT)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGTERM)
	<-sigCh

	go func() {
		<-sigCh
		log.Print("Forced exit")
		os.Exit(1)
	}()
}
