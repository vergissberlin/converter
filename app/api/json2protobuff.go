package main

import (
	"io/ioutil"
	"net/http"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

func convertJSON2ProtobuffHandler(w http.ResponseWriter, r *http.Request) {

	// POST requests only
	if r.Method != http.MethodPost {
		// Set the response status code to 405 (Method Not Allowed)
		w.WriteHeader(http.StatusMethodNotAllowed)

		// Set the response body to "Only POST requests are allowed"
		http.Error(w, "Only POST requests are allowed", http.StatusBadRequest)

		return
	}

	// Read JSON data from the request body
	jsonData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}

	// Convert JSON to Protobuf
	person := &Person{}
	if err := jsonpb.UnmarshalString(string(jsonData), person); err != nil {
		http.Error(w, "Error converting JSON to Protobuf", http.StatusBadRequest)
		return
	}

	// Convert Protobuf data to binary format
	protobufData, err := proto.Marshal(person)
	if err != nil {
		http.Error(w, "Error converting Protobuf to binary format", http.StatusInternalServerError)
		return
	}

	// Send binary Protobuf data as the response
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(protobufData)
}
