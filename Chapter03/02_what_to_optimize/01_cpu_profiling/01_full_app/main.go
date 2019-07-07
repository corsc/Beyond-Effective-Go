package main

import (
	"github.com/PacktPublishing/Advanced-Go-Programming/Chapter03/02_what_to_optimize/01_cpu_profiling/shared"
	"log"
	"net/http"
	"net/http/pprof"
	_ "net/http/pprof"
	"os"
	"os/signal"
)

func main() {
	router := http.NewServeMux()

	// add profiling to our mux
	router.HandleFunc("/debug/pprof/", pprof.Index)
	router.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	router.HandleFunc("/debug/pprof/profile", pprof.Profile)
	router.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	router.HandleFunc("/debug/pprof/trace", pprof.Trace)

	// add our handler
	router.HandleFunc("/", shared.CardShuffler)

	// start the default mux to host the profiling
	go func() {
		log.Println(http.ListenAndServe("0.0.0.0:6060", nil))
	}()

	// start a server
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		_ = server.ListenAndServe()
	}()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)
	<-signalCh
	_ = server.Close()
}
