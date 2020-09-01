//go:generate statik -src=./universal_spectrum_explorer  -include=*.jpg,*.txt,*.html,*.css,*.js
//go:generate goversioninfo -icon=resources/kusterlab.ico

package main

import (
	"bytes"
	"fmt"
	"github.com/bakins/logrus-middleware"
	_ "github.com/kusterlab/use_embedded/statik" // TODO: Replace with the absolute import path
	"github.com/rakyll/statik/fs"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func handlerPeptideAtlas(w http.ResponseWriter, req *http.Request) {
	// we need to buffer the body if we want to read it here and send it
	// in the request.
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// you can reassign the body if you need to parse it as multipart
	req.Body = ioutil.NopCloser(bytes.NewReader(body))

	query := req.URL.Query()
	usi, present := query["usi"] //filters=["color", "price", "brand"]
	if !present || len(usi) == 0 {
		fmt.Println("usi not present")
	}

	// create a new url from the raw RequestURI sent by the client
	// url := fmt.Sprintf("%s://%s%s", proxyScheme, proxyHost, req.RequestURI)
	url := "https://db.systemsbiology.net/sbeams/cgi/PeptideAtlas/ShowObservedSpectrum?USI=" + usi[0]

	fmt.Println(url)

	proxyReq, err := http.NewRequest(req.Method, url, bytes.NewReader(body))

	// We may want to filter some headers, otherwise we could just use a shallow copy
	// proxyReq.Header = req.Header
	proxyReq.Header = make(http.Header)
	for h, val := range req.Header {
		proxyReq.Header[h] = val
	}
	client := &http.Client{
		Timeout: 15 * time.Second}
	resp, err := client.Do(proxyReq)
	fmt.Println(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, err = w.Write(body)
	if err != nil {
		logrus.Info(err)
	}

	// legacy code
}

func handlerJPOST(w http.ResponseWriter, req *http.Request) {
	// we need to buffer the body if we want to read it here and send it
	// in the request.
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	query := req.URL.Query()
	usi, present := query["usi"] //filters=["color", "price", "brand"]
	if !present || len(usi) == 0 {
		fmt.Println("usi not present")
	}

	// you can reassign the body if you need to parse it as multipart
	req.Body = ioutil.NopCloser(bytes.NewReader(body))

	// create a new url from the raw RequestURI sent by the client
	// url := fmt.Sprintf("%s://%s%s", proxyScheme, proxyHost, req.RequestURI)
	// url := "https://repository.jpostdb.org/spectrum/get_data.php?usi=mzspec:PXD005175:CRC_iTRAQ_06:scan:11803:VEYTLGEESEAPGQR/3"
	url := "https://repository.jpostdb.org/spectrum/get_data.php?usi=" + usi[0]
	fmt.Println(url)

	proxyReq, err := http.NewRequest(req.Method, url, bytes.NewReader(body))

	// We may want to filter some headers, otherwise we could just use a shallow copy
	// proxyReq.Header = req.Header
	proxyReq.Header = make(http.Header)
	for h, val := range req.Header {
		proxyReq.Header[h] = val
	}
	client := &http.Client{}
	resp, err := client.Do(proxyReq)
	fmt.Println(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	//	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Content-Encoding", "gzip")

	_, err = w.Write(body)
	if err != nil {
		logrus.Info(err)
	}

	// legacy code
}

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
		Name:   "use",
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
	http.HandleFunc("/peptideAtlas", handlerPeptideAtlas) // Update this line of code
	http.HandleFunc("/jPOST", handlerJPOST)               // Update this line of code
	http.Handle("/use/", l.Handler(http.StripPrefix("/use/", http.FileServer(statikFS)), "use"))

	go http.ListenAndServe(port_string, nil)
	url := "http://localhost" + port_string + "/use/UniversalSpectrumExplorer.html"
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
