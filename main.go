package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/sravzpublic/sravz_golang_crash_course/pkg/aws"
	"github.com/sravzpublic/sravz_golang_crash_course/pkg/util"
)

func main() {
	var wait time.Duration
	var symbols string
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.StringVar(&symbols, "symbols", "ETHBTC,BTCUSDC", "Comma separted symbols list")
	flag.Parse()

	Cryptos = make(map[string]Crypto)
	log.Println("Symbols supported: ", symbols)
	Symbols = util.SplitSting(symbols, ",")
	// Symbols = strings.Split(symbols, ",") // []string{"ETHBTC", "BTCUSDC"}
	wait = util.GetWaitTime()
	log.Println(aws.HelloWord())
	// log.Panicln(aws.HelloWordCannotImport)
	r := mux.NewRouter()
	r.HandleFunc("/currency/{symbol}", GetCrypto).Methods("GET")

	srv := &http.Server{
		Addr: "0.0.0.0:9999",
		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		log.Println("Server listening at localhost:8080")
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	// Blocks main until a signal is received
	<-c
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	// Called after the main function returns
	// Cancel the conext even if the main fuction returns before the timeout of 15 seconds
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)
}
