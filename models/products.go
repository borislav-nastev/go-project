package models

type Product struct {
	ProductId string `json:"productId"`
	ProductName string `json:"productName"`
	ProductDescription string `json:"productDescription"`
	ProductPrice int `json:"price"`
}
