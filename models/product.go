package models

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

type Product struct {
	id          uint    `json:"id"`
	price       float32 `json:"price"`
	image       string  `json:"image"`
	brand       string  `json:"brand"`
	title       string  `json:"title"`
	reviewScore float64 `json:"reviewScore"`
}

func GetById(id int) *Product {

	entity := &Product{}
	productId := url.QueryEscape(strconv.Itoa(id))

	url := fmt.Sprintf("http://challenge-api.luizalabs.com/api/product/%s", productId)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return entity
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return entity
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(entity); err != nil {
		log.Println(err)
	}

	return entity

}
