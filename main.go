//go:generate statik -src=./isv -include=*.jpg,*.txt,*.html,*.css,*.js
//go:generate goversioninfo -icon=resources/icon.ico

package main

import (
	"github.com/bakins/logrus-middleware"
	"github.com/rakyll/statik/fs"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	_ "isv_embed/statik" // TODO: Replace with the absolute import path
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

// ...

func main() {

	var port *int = pflag.Int("port", 8081, "port used for local deployment")
	var verbose *bool = pflag.BoolP("verbose", "v", false, "verbose output")

	var port_string string = ":" + strconv.FormatInt(int64(*port), 10)

	pflag.Parse()
	// viper.BindPFlags(pflag.CommandLine)

	logger := logrus.New()
	if *verbose {
		logger.Level = logrus.InfoLevel
	} else {
		logger.Level = logrus.PanicLevel
	}
	//	logger.Formatter = &logrus.JSONFormatter{}

	l := logrusmiddleware.Middleware{
		Name:   "isv",
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

	go http.ListenAndServe(port_string, nil)
	url := "http://localhost" + port_string + "/isv/PeptideAnnotator.html"
	_ = openURL(url)
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
