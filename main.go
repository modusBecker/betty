package main

import (
	"flag"
	"fmt"

	"github.com/teamjorge/betty/api"
)

func main() {
	var host = flag.String("host", "127.0.0.1", "Host for orchestrator web interface")
	var port = flag.Int("port", 7000, "Port for orchestrator web interface")
	flag.Parse()
	fmt.Printf("Starting API on http://%s:%d\n", *host, *port)
	api.StartAPI(*host, *port)
}
