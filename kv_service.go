package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type AddOrUpdateData struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type DeleteData struct {
	Key string `json:"key"`
}

type KVService struct {
	kvStore map[string]string
}

func (k *KVService) InitKVService() {
	// Initialize in-memory KV store
	k.kvStore = make(map[string]string)

	// Set up routes for API access
	http.HandleFunc("/kv", k.RouteRequests)
}

func (k *KVService) PutResponse(w http.ResponseWriter, req *http.Request) {
	// Get data from request
	reqBody, readErr := io.ReadAll(req.Body)
	if readErr != nil {
		http.Error(w, `{"error": "Could not read request body"}`, http.StatusBadRequest)
		return
	}

	// unmarshal json data
	var putData AddOrUpdateData
	jsonUnmarshalErr := json.Unmarshal(reqBody, &putData)
	if jsonUnmarshalErr != nil {
		http.Error(w, `{"error": "Could not unmarshal request body"}`, http.StatusBadRequest)
		return
	}

	// Check that our post has key and value
	if putData.Key == "" || putData.Value == "" {
		http.Error(w, `{"error": "Key and value must be provided}`, http.StatusBadRequest)
		return
	}

	// Check if value exists. If it does, update it, if not, error
	if _, ok := k.kvStore[putData.Key]; ok {
		k.kvStore[putData.Key] = putData.Value
		resp := make(map[string]string)
		resp[putData.Key] = putData.Value
		jsonResp, jsonMarshalErr := json.Marshal(resp)
		if jsonMarshalErr != nil {
			http.Error(w, `{"error": "Error marshalling response"}`, http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, writeErr := w.Write(jsonResp)
		if writeErr != nil {
			http.Error(w, `{"error": "Error writing response"}`, http.StatusInternalServerError)
		}
		return
	} else {
		// item does not exist, do not update
		http.Error(w, `{"error":"Key Does not exist; Use a POST to create it"}`, http.StatusBadRequest)
		return
	}
}

func (k *KVService) PostResponse(w http.ResponseWriter, req *http.Request) {
	// Get data from request
	reqBody, readErr := io.ReadAll(req.Body)
	if readErr != nil {
		http.Error(w, `{"error": "Could not read request body"}`, http.StatusBadRequest)
		return
	}

	// unmarshal json data
	var postData AddOrUpdateData
	jsonUnmarshalErr := json.Unmarshal(reqBody, &postData)
	if jsonUnmarshalErr != nil {
		http.Error(w, `{"error": "Could not unmarshal request body"}`, http.StatusBadRequest)
		return
	}

	// Check that our post has key and value
	if postData.Key == "" || postData.Value == "" {
		http.Error(w, `{"error": "Key and value must be provided}`, http.StatusBadRequest)
		return
	}

	// Check if value exists. If it does not, insert it. If it does, error. Updates to a key should be done via a PUT operation
	if _, ok := k.kvStore[postData.Key]; !ok {
		// Item is not present, insert it
		k.kvStore[postData.Key] = postData.Value

		// return json denoting everything is ok
		resp := make(map[string]string)
		resp[postData.Key] = postData.Value
		jsonResp, jsonMarshalError := json.Marshal(resp)
		if jsonMarshalError != nil {
			http.Error(w, `{"error": "Error marshaling response"}`, http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		_, writeErr := w.Write(jsonResp)
		if writeErr != nil {
			http.Error(w, `{"error": "Error writing response"}`, http.StatusInternalServerError)
		}
		return
	} else {
		// Item exits, error out
		http.Error(w, `{"error": "Key `+postData.Key+` already exists; Use a PUT to update it"}`, http.StatusInternalServerError)
		return
	}
}

func (k *KVService) GetResponse(w http.ResponseWriter, req *http.Request) {
	// get data from request
	// TODO: update to read from GET body
	key := req.URL.Query().Get("key")
	// Check if we have the key in our store.
	if val, ok := k.kvStore[key]; !ok {
		// key is not present, return error
		http.Error(w, `{"error": "Key not present"}`, http.StatusInternalServerError)
	} else {
		//return key and value in json format
		resp := make(map[string]string)
		resp[key] = val
		jsonResp, jsonMarshalError := json.Marshal(resp)
		if jsonMarshalError != nil {
			http.Error(w, `{"error": "Error marshaling response"}`, http.StatusInternalServerError)
			return
		}

		_, writeErr := w.Write(jsonResp)
		if writeErr != nil {
			http.Error(w, `{"error": "Error writing response"}`, http.StatusInternalServerError)
		}
		return
	}
}

func (k *KVService) DeleteResponse(w http.ResponseWriter, req *http.Request) {
	// delete a value from kvStore

	reqBody, readErr := io.ReadAll(req.Body)
	if readErr != nil {
		http.Error(w, `{"error": "Could not read request body"}`, http.StatusBadRequest)
		return
	}

	// Pull key out of req body
	var deleteData DeleteData
	jsonUnmarshalErr := json.Unmarshal(reqBody, &deleteData)
	if jsonUnmarshalErr != nil {
		http.Error(w, `{"error": "Could not unmarshal request body"}`, http.StatusBadRequest)
		return
	}

	// Check if we have the key in our store.
	if _, ok := k.kvStore[deleteData.Key]; !ok {
		// key is not present, return error
		http.Error(w, `{"error": "Key not present"}`, http.StatusInternalServerError)
		return
	} else {
		// delete and send response
		delete(k.kvStore, deleteData.Key)
		w.WriteHeader(http.StatusOK)
		_, writeErr := w.Write([]byte(`{"message": "Deleted ` + deleteData.Key + `"}`))
		if writeErr != nil {
			http.Error(w, `{"error": "Error writing response"}`, http.StatusInternalServerError)
			return
		}
		return
	}
}
func (k *KVService) RouteRequests(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		// POST adds a new KV to our store, and errors if there is already an existing key.
		k.PostResponse(w, req)
	}

	if req.Method == "PUT" {
		// Update value
		k.PutResponse(w, req)
	}

	if req.Method == "GET" {
		// Get Value
		k.GetResponse(w, req)
	}

	if req.Method == "DELETE" {
		// Delete value from store
		k.DeleteResponse(w, req)
	}
}
