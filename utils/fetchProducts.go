package utils

import (
	"bankapp/models"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"math/rand/v2"
	"fmt"
)

func FetchProducts() ([]models.Product, error){
	resp, err := http.Get("https://fakestoreapi.com/products")

	if err != nil {
		log.Fatalf("Erro ao fazer request: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("API request falhou com status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Erro ao ler resposta do body: %v", err)
	}

	var products []models.Product

	err = json.Unmarshal(body, &products)
	if err != nil {
		log.Fatalf("Erro ao decodificar JSON: %v", err)
	}

	fmt.Printf("Products: %v", products[:5])

	return products, nil 
}

func ShufflePurchase(products []models.Product) models.Product{
	i := rand.IntN(len(products))
	product := products[i]

	return product
}
