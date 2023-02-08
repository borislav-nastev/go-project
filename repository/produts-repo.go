package repository

import (
	"context"
	"go/models"
	"log"

	"cloud.google.com/go/firestore"
)

type ProductRepository interface {
	Save(product *models.Product) (*models.Product, error)
	FindAll() ([]models.Product, error)
}

type repo struct{}

// NewProductRepository
func NewProductRepository() ProductRepository {
	return &repo{}
}

const (
	projectId string = "go-project-26def"
	collectionName string = "products"
)

func (*repo) Save(product *models.Product) (*models.Product, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)

	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}

	defer client.Close()

	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ProductId": product.ProductId,
		"ProductName": product.ProductName,
		"ProductDescription": product.ProductDescription,
		"ProductPrice": product.ProductPrice,
	})

	if err != nil {
		log.Fatalf("Failed adding a new product: %v", err)
		return nil, err
	}

	return product, nil
}

func (*repo) FindAll() ([]models.Product, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)

	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}

	defer client.Close()

	var products []models.Product
	iterator := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := iterator.Next()
		if err != nil {
			log.Fatalf("Failed to iterate the list of products: %v", err)
			return nil, err
		}
		product := models.Product{
			ProductId: doc.Data()["ProductId"].(string),
			ProductName: doc.Data()["ProductName"].(string),
			ProductDescription: doc.Data()["ProductDescription"].(string),
			ProductPrice: doc.Data()["ProductPrice"].(int),
		}
		products = append(products, product)
	}

	return products, nil
}