package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Product struct {
	ProductID		int		`json:"productId"`
	Manufacturer	string	`json:"manufacturer"`
	Sku				string	`json:"sku"`
	Upc				string	`json:"upc"`
	PricePerUnit	string	`json:"pricePerUnit"`
	QualityOnHand	int		`json:"quantityOnHand"`
	ProductName		string	`json:"productName"`
}

var products []Product

func productsHandler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		bs, err := json.Marshal(products)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(bs)
	case http.MethodPost:
		bs, err := ioutil.ReadAll(request.Body)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		var p Product
		err = json.Unmarshal(bs, &p)
		if err != nil {
			writer.WriteHeader(http.StatusBadRequest)
			return
		}

		products = append(products, p)
		writer.WriteHeader(http.StatusCreated)
	}
}

func productHandler(writer http.ResponseWriter, request *http.Request) {
	urlPathSegments := strings.Split(request.URL.Path, "products/")
	productID, err := strconv.Atoi(urlPathSegments[len(urlPathSegments) - 1])
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	for _, p := range products {
		if p.ProductID == productID {
			bs, err := json.Marshal(p)
			if err != nil {
				writer.WriteHeader(http.StatusInternalServerError)
				return
			}
			writer.Header().Set("Content-Type", "application/json")
			writer.Write(bs)
			return
		}
	}
	writer.WriteHeader(http.StatusNotFound)
	return
}

func main() {
	http.HandleFunc("/products", productsHandler)
	http.HandleFunc("/product", productHandler)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
