package main

import (
	"flag"
	"github.com/qinix/portdev"
)

var (
	port     int
)

func main() {
	flag.IntVar(&port, "port", 8080, "Port to listen")
	flag.Parse()

	portdev.Listen(port)
}
