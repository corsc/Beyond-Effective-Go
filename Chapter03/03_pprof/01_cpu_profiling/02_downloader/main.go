package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"runtime/pprof"
)

func main() {
	// add profiling
	f, err := os.Create("cpu.pprof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	// main code
	connection := makeConnection()
	data := download(connection)
	closeConnection(connection)

	fmt.Printf("Download was: %v", data)
}

func makeConnection() net.Conn {
	waste := 0
	for x := 0; x < 10*100000000; x++ {
		// waste some CPU
		waste++
	}
	println(waste)

	// not implemented
	return nil
}

func download(conn net.Conn) []byte {
	waste := 0
	for x := 0; x < 100*100000000; x++ {
		// waste some CPU
		waste++
	}
	println(waste)

	// not implemented
	return nil
}

func closeConnection(conn net.Conn) {
	waste := 0
	for x := 0; x < 5*100000000; x++ {
		// waste some CPU
		waste++
	}
	println(waste)

	// not implemented
}
