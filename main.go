//go:generate statik -src=./isv -include=*.jpg,*.txt,*.html,*.css,*.js
//go:generate goversioninfo -icon=resources/icon.ico -gofile=./resources/versioninfo.json

package main

import (
	"github.com/bakins/logrus-middleware"
	"github.com/rakyll/statik/fs"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "isv_embed/statik" // TODO: Replace with the absolute import path
)

// ...

func main() {

	pflag.Int("flagname", 1234, "help message for flagname")

	pflag.Parse()
	// viper.BindPFlags(pflag.CommandLine)

	logger := logrus.New()
	logger.Level = logrus.InfoLevel
	logger.Formatter = &logrus.JSONFormatter{}

	l := logrusmiddleware.Middleware{
		Name:   "example",
		Logger: logger,
	}

	statikFS, err := fs.New()
	if err != nil {
		logger.Fatal(err)
	}

	// Serve the contents over HTTP.
	//	http.Handle("/isv/", http.StripPrefix("/isv/", http.FileServer(statikFS)))
	//http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(statikFS)))
	// http.Handle("/", l.Handler(http.HandlerFunc(handler), "homepage"))
	http.Handle("/isv/", l.Handler(http.StripPrefix("/isv/", http.FileServer(statikFS)), "isv"))
	logger.Info("test")

	go http.ListenAndServe(":8081", nil)
	const url = "http://localhost:8081/isv/PeptideAnnotator.html"
	openURL(url)
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	// `signal.Notify` registers the given channel to
	// receive notifications of the specified signals.
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// This goroutine executes a blocking receive for
	// signals. When it gets one it'll print it out
	// and then notify the program that it can finish.
	go func() {
		sig := <-sigs
		logger.Println()
		logger.Println(sig)
		done <- true
	}()

	// The program will wait here until it gets the
	// expected signal (as indicated by the goroutine
	// above sending a value on `done`) and then exit.
	logger.Println("awaiting signal")
	<-done
	logger.Println("exiting")
}
