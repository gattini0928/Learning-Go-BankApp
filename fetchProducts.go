package main

import (
	"bankapp/models"
	"encoding/json"
	"io"
	"fmt"
	"log"
	"net/http"
)

var products []models.Product

func fetchProducts() {
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

	err = json.Unmarshal(body, &products)
	if err != nil {
		log.Fatalf("Erro ao decodificar JSON: %v", err)
	}

	fmt.Println("Titulo", products[0].Title)
	fmt.Println("preço", products[0].Price)
}