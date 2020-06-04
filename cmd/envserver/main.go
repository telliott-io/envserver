package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/", env)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func env(w http.ResponseWriter, req *http.Request) {
	var allVars = make(map[string]string)
	for _, element := range os.Environ() {
		variable := strings.Split(element, "=")
		// Add the environment var to the output map, excluding all
		// Kubernetes-specific environment variables (which expose API details)
		if !strings.HasPrefix(variable[0], "KUBERNETES_") {
			allVars[variable[0]] = variable[1]
		}
	}
	j, err := json.MarshalIndent(allVars, "", " ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, string(j))
}
