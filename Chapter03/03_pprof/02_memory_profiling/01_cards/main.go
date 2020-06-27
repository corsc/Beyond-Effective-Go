package main

import (
	"log"
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"

	"github.com/corsc/Advanced-Go-Programming/Chapter03/03_pprof/01_cpu_profiling/game"
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
	router.HandleFunc("/", CardShuffler)

	// start the default mux to host the profiling
	go func() {
		log.Println(http.ListenAndServe("0.0.0.0:6060", nil))
	}()

	// start a server
	server := &http.Server{
		Addr:    ":8888",
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

func CardShuffler(resp http.ResponseWriter, req *http.Request) {
	// create a deck of cards
	cards := game.NewDeck()

	// intentionally shuffle a few times (to spend lots of CPU)
	for x := 0; x < 100; x++ {
		game.Shuffle(cards)
	}

	// return the result
	for index, card := range cards {
		if index > 0 {
			_, _ = resp.Write([]byte(", "))
		}
		_, _ = resp.Write([]byte(card.Face))
		_, _ = resp.Write([]byte(card.Suit))
	}
}
