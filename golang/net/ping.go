package main

import (
	"fmt"

	"github.com/go-ping/ping"
)

func Ping(url string) {
	pinger, err := ping.NewPinger(url)
	if err != nil {
		panic(err)
	}
	pinger.Count = 3
	err = pinger.Run() // Blocks until finished.
	if err != nil {
		panic(err)
	}
	stats := pinger.Statistics() // get send/receive/duplicate/rtt stats
	fmt.Println(stats)
}

// go get -u github.com/go-ping/ping
func main() {
	Ping("127.0.0.1")
}
