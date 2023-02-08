package controllers

import (
	"encoding/json"
	"fmt"
	"go/models"

	//"go/repository"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var products []models.Product
var product []models.Product

// var (
// 	repo repository.ProductRepository = repository.NewProductRepository()
// )

// type Product struct {
// 	ProductId string `json:"productId"`
// 	ProductName string `json:"productName"`
// 	ProductDescription string `json:"productDescription"`
// 	ProductPrice int `json:"price"`
// }

func init() {
	products = []models.Product{
		{ProductId:"1", ProductName: "Product 1", ProductDescription: "Product 1 description", ProductPrice: 234},
		{ProductId:"2", ProductName: "Product 2", ProductDescription: "Product 2 description", ProductPrice: 299},
		{ProductId:"3", ProductName: "Product 3", ProductDescription: "Product 3 description", ProductPrice: 320},
		{ProductId:"4", ProductName: "Product 4", ProductDescription: "Product 4 description", ProductPrice: 240}}
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint hit: GetAllProducts")
	w.Header().Set("Content-Type", "application/json")

	resp := make(map[string]interface{})
	resp["message"] = "Status OK products are here"
	resp["data"] = products

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error marshalling the products array"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
	return
}

func GetSingleProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println(("Endpoint hit: GetSingleProduct"))
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	resp := make(map[string]interface{})
	fmt.Println(params["id"])
	fmt.Println(products)

	for _, product := range products{
		if product.ProductId == params["id"]{
			resp["message"] = "Status Ok product is here"
			resp["data"] = product
			jsonResp, err := json.Marshal(resp)
			if err != nil {
				log.Fatalf("Error happened in JSON marshal. Err: %s", err)
			}

			w.WriteHeader(http.StatusOK)
			w.Write(jsonResp)
			return
		} 
	}

	resp["message"] = "Not found"
	resp["data"] = nil

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	w.WriteHeader(http.StatusNotFound)
	w.Write(jsonResp)
	return
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println(("Endpoint hit: CreateProduct"))
	w.Header().Set("Content-Type", "application/json")

	var product models.Product
	err := json.NewDecoder((r.Body)).Decode(&product)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error unMarshalling the request"}`))
		return
	}

	stringID := strconv.Itoa(len(products) + 1)
	product.ProductId = stringID

	products = append(products, product)

	resp := make(map[string]interface{})
	resp["message"] = "Status OK product was created"
	resp["data"] = product

	jsonResp, err := json.Marshal(resp)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}

func EditProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println(("Endpoint hit: EditProduct"))
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resp := make(map[string]interface{})
	resp["message"] = "Status OK product was edited"
	resp["data"] = product
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
	resp["message"] = "Status OK product was deleted"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	return
}