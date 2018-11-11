package main

import (
	"flag"
	. "time-logger/internal/app/config"
)

func main() {
	portPtr := flag.Int("port", 3000, "Port number")
	flag.Parse()

	StartServer(*portPtr)
}
