package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server interface {
	Start(ctx context.Context) error
	Pause() error
	Resume() error
	Stop() error
	Restart() error
	Ping() error
}
type AppServer struct {
	label       string
	isPause     bool
	address     string
	server      *http.Server
	handlerFunc http.Handler
	stopCn      chan struct{}
	beforeStart func()
}

func (app *AppServer) Start(ctx context.Context) error {
	fmt.Println(app.label, " Server trying to start at ", app.address)
	//check address and handler is nil
	app.server = &http.Server{
		Addr:    app.address,
		Handler: app.handlerFunc,
	}

	go func() {
		<-app.stopCn
		fmt.Printf("%s is closing... \n", app.label)
		if app.server != nil {
			err := app.server.Shutdown(ctx)
			if err != nil {
				log.Printf("server Shutdown: %v", err)
			}
		}
	}()

	return app.server.ListenAndServe()

}
func (app *AppServer) Stop() error {
	fmt.Println(app.label, " Server stop")

	return nil
}
func (app *AppServer) Restart() error {
	fmt.Println(app.label, " Server restart")
	return nil
}
func (app *AppServer) Pause() error {
	fmt.Println(app.label, " Server Pause")
	app.isPause = true
	return nil
}
func (app *AppServer) Resume() error {
	fmt.Println(app.label, " Server Resume")
	app.isPause = false
	return nil
}
func (app *AppServer) Ping() error {
	//check if I am live
	return nil
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	stop := make(chan struct{})
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(3*time.Second))
	defer cancel()
	serverList := []Server{
		&AppServer{
			label:   "Http Server",
			address: ":8080",
			stopCn:  stop,
			handlerFunc: http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					w.Write([]byte("This is Http Server Root"))
				},
			),
		},
		&AppServer{
			label:   "Debug Server",
			address: ":8081",
			stopCn:  stop,
			handlerFunc: http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					w.Write([]byte("This is Debug Server Root"))
				},
			),
		},
		&AppServer{
			label:   "Admin Console",
			address: ":8082",
			stopCn:  stop,
			handlerFunc: http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					w.Write([]byte("This is Admin Server Root"))
				},
			),
		},
	}

	done := make(chan error, len(serverList))
	// start all servers
	for _, app := range serverList {
		go func(server Server) {
			done <- server.Start(ctx)
		}(app)
	}
	// Wait for Control C or Kill to exit
	ch := make(chan os.Signal, 1)
	signals := []os.Signal{syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT}
	signal.Notify(ch, signals...)
	var stopped bool
	go func() {
		<-ch
		stopped = true
		close(stop)
	}()

	for i := 0; i < len(serverList); i++ {

		if err := <-done; err != nil {
			fmt.Printf("error: %v \n", err)
		}

		if !stopped {
			stopped = true
			close(stop)
		}
	}

}
