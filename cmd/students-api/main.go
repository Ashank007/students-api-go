package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Ashank007/students-api-go/internal/config"
)

func main(){

	cfg := config.MustLoad()

	router := http.NewServeMux()

	router.HandleFunc("GET /",func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("Welcome to Students Api"))
  
	})

	server := http.Server {
   Addr: cfg.HTTPServer.Address,
	 Handler: router,
	}

	slog.Info("Server Started ",slog.String("Address",cfg.HTTPServer.Address))
	fmt.Printf("Server Started %s", cfg.HTTPServer.Address)

	done := make(chan os.Signal,1)

	signal.Notify(done,os.Interrupt,syscall.SIGINT,syscall.SIGTERM)

	go func ()  {
			err := server.ListenAndServe()
			if err != nil {
				log.Fatal("Fail to Start Server")
			}
	} ()
	
	<-done

	slog.Info("Shutting Down the Server")


	ctx,cancel := context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Failed to ShutDown Server",slog.String("error",err.Error()))
	}

	slog.Info("Server ShutDown Sucessfully")
}
