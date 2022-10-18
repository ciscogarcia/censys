package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func InitTestService() {
	http.HandleFunc("/test", TestEndpoints)
}

func TestDeleteKV(w http.ResponseWriter, r *http.Request) {
	// Test deleting existing key
	body, _ := json.Marshal(map[string]string{
		"key": "Hello",
	})
	responseBody := bytes.NewBuffer(body)
	req, err := http.NewRequest(http.MethodDelete, "http://localhost:10000/kv", responseBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)
	w.Write(respBody)

	// Test deleting non-existant key
	body, _ = json.Marshal(map[string]string{
		"key": "Cat",
	})
	responseBody = bytes.NewBuffer(body)
	req, err = http.NewRequest(http.MethodDelete, "http://localhost:10000/kv", responseBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp, err = httpClient.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	respBody, _ = io.ReadAll(resp.Body)
	w.Write(respBody)
}

func TestGetKV(w http.ResponseWriter, r *http.Request) {
	// Test getting a key that exists
	resp, _ := http.Get("http://localhost:10000/kv?key=Hello")
	respBody, _ := io.ReadAll(resp.Body)
	w.Write(respBody)

	// Test getting a key that does not exist
	resp, _ = http.Get("http://localhost:10000/kv?key=linux")
	respBody, _ = io.ReadAll(resp.Body)
	w.Write(respBody)
}

func TestUpdateKV(w http.ResponseWriter, r *http.Request) {
	// Test updating "Hello" -> "Censys"
	body, _ := json.Marshal(map[string]string{
		"key":   "Hello",
		"value": "Censys",
	})
	responseBody := bytes.NewBuffer(body)
	req, err := http.NewRequest(http.MethodPut, "http://localhost:10000/kv", responseBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)
	w.Write(respBody)

	// Test updating "Hello" -> "World"
	body, _ = json.Marshal(map[string]string{
		"key":   "Hello",
		"value": "World",
	})
	responseBody = bytes.NewBuffer(body)
	req, err = http.NewRequest(http.MethodPut, "http://localhost:10000/kv", responseBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp, err = httpClient.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	respBody, _ = io.ReadAll(resp.Body)
	w.Write(respBody)

	// Test invalid data
	body, _ = json.Marshal(map[string]string{
		"key": "Hello",
	})
	responseBody = bytes.NewBuffer(body)
	req, err = http.NewRequest(http.MethodPut, "http://localhost:10000/kv", responseBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resp, err = httpClient.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	respBody, _ = io.ReadAll(resp.Body)
	w.Write(respBody)

}

func TestAddKV(w http.ResponseWriter, r *http.Request) {
	// Test adding "Hello" -> "World"
	body, _ := json.Marshal(map[string]string{
		"key":   "Hello",
		"value": "World",
	})
	responseBody := bytes.NewBuffer(body)
	resp, err := http.Post("http://localhost:10000/kv", "application/json", responseBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)
	w.Write(respBody)

	// Test Adding a Key that already exists
	responseBody = bytes.NewBuffer(body)
	resp, err = http.Post("http://localhost:10000/kv", "application/json", responseBody)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	respBody, _ = io.ReadAll(resp.Body)
	w.Write(respBody)

	// Test adding invalid KV
	body, _ = json.Marshal(map[string]string{
		"key": "Hello",
	})
	responseBody = bytes.NewBuffer(body)
	resp, err = http.Post("http://localhost:10000/kv", "application/json", responseBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	respBody, _ = io.ReadAll(resp.Body)
	w.Write(respBody)
}

func TestEndpoints(w http.ResponseWriter, r *http.Request) {
	TestAddKV(w, r)
	TestUpdateKV(w, r)
	TestGetKV(w, r)
	TestDeleteKV(w, r)
}
