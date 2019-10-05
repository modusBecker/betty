package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/teamjorge/betty/orchestration"
	"github.com/teamjorge/betty/structures"
)

// handlerGetContainers handles requests for the /get_containers endpoint
func handlerGetContainers(w http.ResponseWriter, r *http.Request) {
	containers := orchestration.GetRunningContainers()
	if containers == nil {
		fmt.Fprintf(w, "[]")
		return
	}
	containerJSON, err := json.Marshal(containers)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, string(containerJSON))
}

// handlerStartContainer handles requests for the /start_container endpoint
func handlerStartContainer(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var requestData structures.StartContainerRequest
	err = json.Unmarshal(body, &requestData)
	if err != nil {
		panic(err)
	}
	resp := orchestration.StartContainer(requestData)
	response := structures.StartContainerResponse{
		ID:       resp.ID,
		Warnings: resp.Warnings,
	}
	responseJSON, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}
	w.Write(responseJSON)
}

// handlerGetContainerByName handles request for /get_container_by_name
func handlerGetContainerByName(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	containerName := queryParams.Get("name")
	// Handle empty query parameters
	if containerName == "" {
		errorResponse := structures.ErrorResponse{
			Error:   true,
			Message: "name query parameter required",
			Detail:  "",
		}
		errorResponseJSON, err := json.Marshal(errorResponse)
		if err != nil {
			panic(err)
		}
		w.WriteHeader(400)
		w.Write(errorResponseJSON)
		return
	}

	res := orchestration.GetContainerByName(containerName)
	if len(res) == 0 {
		fmt.Fprintf(w, "[]")
		return
	}
	resJSON, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}
	w.Write(resJSON)
}

// handerListNetwork handles requests for the /list_networks endpoint
func handlerListNetwork(w http.ResponseWriter, r *http.Request) {
	networks := orchestration.ListNetworks()
	networksJSON, err := json.Marshal(networks)
	if err != nil {
		panic(err)
	}
	w.Write(networksJSON)
}
