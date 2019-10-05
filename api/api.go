package api

import (
	"fmt"
	"log"
	"net/http"
)

// StartAPI adds handlers and starts up the API server
func StartAPI(host string, port int) {
	http.HandleFunc("/get_containers", handlerGetContainers)
	http.HandleFunc("/get_container_by_name", handlerGetContainerByName)
	http.HandleFunc("/start_container", handlerStartContainer)
	http.HandleFunc("/list_networks", handlerListNetwork)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), nil))
}
