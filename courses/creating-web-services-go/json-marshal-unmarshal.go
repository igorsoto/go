package main

import (
	"encoding/json"
	"fmt"
	"log"
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

func marshalProduct(p *Product) (string, error) {
	pJSON, err := json.Marshal(p)
	if err != nil {
		return "", err
	}
	return string(pJSON), nil
}

func unmarshalProduct(j string) (*Product, error) {
	p := &Product{}
	err := json.Unmarshal([]byte(j), p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func main() {
	p := &Product{
		ProductID: 123,
		Manufacturer: "Big Box Company",
		Sku: "4561qHJK",
		Upc: "414654444566",
		PricePerUnit: "12.99",
		QualityOnHand: 28,
		ProductName: "Gizmo",
	}

	j, err := marshalProduct(p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Product marshal:", j)

	j = strings.Replace(j, "Big Box Company", "Small Box Company", -2)
	p, err = unmarshalProduct(j)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Product unmarshall:", p.Manufacturer)
}
