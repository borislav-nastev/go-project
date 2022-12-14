package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)


func GetAllProducts(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint hit: GetAllProducts")
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "Status OK -> products are here"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	return
}

func GetSingleProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println(("Endpoint hit: GetSingleProduct"))
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "Status Ok -> product is here"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	return
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println(("Endpoint hit: CreateProduct"))
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "Status OK -> product was created"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	return
}

func EditProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println(("Endpoint hit: EditProduct"))
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "Status OK -> product was edited"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	return
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println(("Endpoint hit: DeleteProduct"))
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]string)
	resp["message"] = "Status OK -> product was deleted"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	return
}