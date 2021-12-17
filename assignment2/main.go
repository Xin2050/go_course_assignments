package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
}
type AppServer struct {
	label       string
	address     string
	server      *http.Server
	handlerFunc http.Handler
	beforeStart func()
}

func (app *AppServer) Stop(ctx context.Context) error {
	fmt.Printf("%s is closing...\n", app.label)
	return app.server.Shutdown(ctx)
}
func (app *AppServer) Start(ctx context.Context) error {
	fmt.Println(app.label, " Server trying to start at ", app.address)
	//check address and handler is nil
	app.server = &http.Server{
		Addr:    app.address,
		Handler: app.handlerFunc,
	}
	if app.beforeStart != nil {
		app.beforeStart()
	}

	err := app.server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func main() {

	serverList := []Server{
		&AppServer{
			label:   "Http Server",
			address: ":8080",
			handlerFunc: http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					w.Write([]byte("This is Http Server Root"))
				},
			),
			beforeStart: func() {
				time.Sleep(time.Millisecond * 100)
			},
		},
		&AppServer{
			label:   "Debug Server",
			address: ":8081",
			handlerFunc: http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					w.Write([]byte("This is Debug Server Root"))
				},
			),
			beforeStart: func() {
				time.Sleep(time.Millisecond * 100)
			},
		},
		&AppServer{
			label:   "Admin Console",
			address: ":8083", // if you change port to 8081, this will case all server down
			handlerFunc: http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					w.Write([]byte("This is Admin Server Root"))
				},
			),
			beforeStart: func() {
				fmt.Println("Run some long time compute code")
				time.Sleep(time.Millisecond * 100)
			},
		},
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	group, ctx := errgroup.WithContext(ctx)
	//start all the servers
	for _, server := range serverList {
		server := server
		group.Go(func() error {
			return server.Start(ctx)
		})
		group.Go(func() error {
			select {
			case <-ctx.Done():
				// use below line will activate exit timeout
				//time.Sleep(time.Second * 4)
				server.Stop(ctx)
				return ctx.Err()
			}
		})
	}
	// Wait for Control C or Kill to exit
	signals := []os.Signal{syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGKILL}
	cn := make(chan os.Signal, 1)
	signal.Notify(cn, signals...)
	group.Go(func() error {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-cn:
			fmt.Println("\n received an exit command, start to try closing....")
			cancel()
			//set up a timeout
			time.AfterFunc(time.Second*3, func() {
				fmt.Println("Exit process time out. I will force to exit the process")
				os.Exit(1)
			})
		}
		return nil
	})
	if err := group.Wait(); err != nil {
		fmt.Printf("Error :%v", err)
	}

}
