package main

import (
	"flag"
	"fmt"
	"github.com/yusufnuru/femProject/internal/app"
	"github.com/yusufnuru/femProject/internal/routes"
	"net/http"
	"time"
)

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "go backend server port")
	flag.Parse()

	newApp, err := app.NewApp()
	if err != nil {
		panic(err)
	}

	r := routes.SetupRoutes(newApp)
	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", port),
		Handler:           r,
		IdleTimeout:       time.Minute,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      30 * time.Second,
	}

	newApp.Logger.Printf("We are running on port %d", port)
	err = server.ListenAndServe()
	if err != nil {
		newApp.Logger.Fatal(err)
	}

}
