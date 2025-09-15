package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/quii/learn-go-with-tests/kvstore"
)

// A struct to hold the key-value store and its related handlers
type server struct {
	store *kvstore.KVStore
}

type ResponseMessage struct {
	Message string `json:"message"`
}

type GetPayload struct {
	Key string `json:"key"`
}

type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (s *server) GetHandler(w http.ResponseWriter, r *http.Request) {
	var p GetPayload

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	if err := decoder.Decode(&p); err != nil {
		fmt.Println(err)
		return
	}

	key := p.Key

	value, error := s.store.Get(key)
	if error != nil {
		http.Error(w, "Key not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(KeyValue{key, value})
}

func (s *server) SetHandler(w http.ResponseWriter, r *http.Request) {
	var p KeyValue
	w.Header().Set("Content-Type", "application/json")

	error := json.NewDecoder(r.Body).Decode(&p)
	if error != nil {
		json.NewEncoder(w).Encode(ResponseMessage{"cannot decode request body"})
		return
	}

	if p.Key == "" || p.Value == "" {
		json.NewEncoder(w).Encode(ResponseMessage{"cannot have empty values"})
		return
	}

	s.store.Set(p.Key, p.Value)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ResponseMessage{"value updated"})
}

func (s *server) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	var p KeyValue
	w.Header().Set("Content-Type", "application/json")

	error := json.NewDecoder(r.Body).Decode(&p)
	if error != nil {
		json.NewEncoder(w).Encode(ResponseMessage{"cannot decode request body"})
		return
	}

	if p.Key == "" || p.Value == "" {
		json.NewEncoder(w).Encode(ResponseMessage{"cannot have empty values"})
		return
	}

	_, error = s.store.Delete(p.Key)
	if error != nil {
		json.NewEncoder(w).Encode(ResponseMessage{"key does not exist"})
		return
	}
	json.NewEncoder(w).Encode(ResponseMessage{"value updated"})
}

func main() {
	s := &server{
		store: kvstore.NewKVStore(),
	}

	router := chi.NewRouter()
	//router.Use(middleware.Logger)

	router.Get("/get", s.GetHandler)
	router.Post("/set", s.SetHandler)
	router.Post("/delete", s.DeleteHandler)

	fmt.Println("Server listening on :8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
