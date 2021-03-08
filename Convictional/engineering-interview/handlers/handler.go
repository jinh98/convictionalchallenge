package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinh98/convictional/engineering-interview/model"
)

var ser *model.Service

// GetProducts retrieves products from given url
func GetProducts() ([]model.Product, error) {
	resp, err := http.Get("https://my-json-server.typicode.com/convictional/engineering-interview/products")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer resp.Body.Close()
	// parse json into products
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	// fmt.Println(string(bytes))
	var products []model.Product

	ser, err = model.NewService()
	if err != nil {
		log.Println(err)
	}
	err = json.Unmarshal(bytes, &products)
	if err != nil {
		log.Println(err)
	}
	for i := range products {
		ser.AddProduct(&products[i])
	}

	return products, err
}

// FindProductById writes a response with the product id as a string
func FindProductById(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	w.Header().Set("Content-Type", "text/plain")
	str, _ := json.Marshal(ser.GetProduct(id))
	w.Write([]byte(str))
}
