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
	"github.com/Ashank007/students-api-go/internal/http/handlers/student"
	"github.com/Ashank007/students-api-go/internal/storage/sqlite"
)

func main(){

	cfg := config.MustLoad()
	
	storage,err := sqlite.New(cfg)

	if err != nil {
		log.Fatal(err)
	}

	slog.Info("Storage Initialized",slog.String("env",cfg.Env),slog.String("version","1.0.0"))

	router := http.NewServeMux()

	router.HandleFunc("POST /api/students",student.New(storage))

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
