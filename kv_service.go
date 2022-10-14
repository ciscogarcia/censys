package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type PostData struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type GetData struct {
	Key string `json:"key"`
}

type KVService struct {
	kvStore map[string]string
}

func (k *KVService) InitKVService() {
	// Initialize in-memory KV store
	k.kvStore = make(map[string]string)

	// Set up routes for API access
	http.HandleFunc("/kv", k.AddKVToStore)
}

func (k *KVService) AddKVToStore(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		// POST adds a new KV to our store, and errors if there is already an existing key.

		// Get data from request
		reqBody, readErr := io.ReadAll(req.Body)
		if readErr != nil {
			http.Error(w, "Could not read request body", http.StatusBadRequest)
			return
		}

		// unmarshal json data
		var postData PostData
		jsonUnmarshalErr := json.Unmarshal(reqBody, &postData)
		if jsonUnmarshalErr != nil {
			http.Error(w, "Could not unmarshal request body", http.StatusBadRequest)
			return
		}

		// Check if value exists. If it does not, insert it. If it does, error. Updates to a key should be done via a PUT operation
		if _, ok := k.kvStore[postData.Key]; !ok {
			// Item is not present, insert it
			k.kvStore[postData.Key] = postData.Value

			// return json denoting everything is ok
			resp := make(map[string]string)
			resp["message"] = fmt.Sprintf("Created key %s with value %s", postData.Key, postData.Value)
			jsonResp, jsonMarshalError := json.Marshal(resp)
			if jsonMarshalError != nil {
				http.Error(w, "Error marshalling response", http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusCreated)
			w.Write(jsonResp)
			return
		} else {
			// Item exits, error out
			http.Error(w, "Key "+postData.Key+" already exists; Use a PUT to update it", http.StatusInternalServerError)
			return
		}

	}
	if req.Method == "PUT" {
		// Update value
	}

	if req.Method == "GET" {
		// Get Value

		// get data from request
		reqBody, readErr := io.ReadAll(req.Body)
		if readErr != nil {
			http.Error(w, "Could not read requestBody", http.StatusBadRequest)
			return
		}

		// Pull key out of req body
		var getData GetData
		jsonUnmarshalErr := json.Unmarshal(reqBody, &getData)
		if jsonUnmarshalErr != nil {
			http.Error(w, "Could not unmarshal request body", http.StatusBadRequest)
			return
		}

		// Check if we have the key in our store.
		if val, ok := k.kvStore[getData.Key]; !ok {
			// key is not present, return error
			http.Error(w, "Key not present", http.StatusInternalServerError)
		} else {
			//return key and value in json format

		}
	}

	if req.Method == "DELETE" {
		// Delete value from store
	}
}
