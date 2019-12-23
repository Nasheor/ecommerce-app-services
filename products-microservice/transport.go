package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

type getProductsRequest struct {}

type getProductsResponse struct {
	I string `json:"productid"`
	N string `json:"name"`
	QUANTITY int64 `json: "quantity"`
	PRICE float64 `json: "price"`
}

type deleteProductsRequest struct {
	ID string `json:"id"`
	Q int64 `json:"quantity"`
	A bool `json:"admin"`
}

type deleteProductResponse struct {
	S bool `json:success`
}

func makeProductsEndpoint(ps ProductsService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		//req := request.(getProductsRequest)
		products := ps.getProducts()
		productsResponse := make([]*getProductsResponse, 0)
		for _, product := range products {
			tmp_holder := new(getProductsResponse)
			tmp_holder.I = product.productID
			tmp_holder.N = product.name
			tmp_holder.PRICE = product.price
			tmp_holder.QUANTITY = product.quantity
			productsResponse = append(productsResponse, tmp_holder)
		}
		return productsResponse, nil
	}
}

func decodeGetProductsRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request getProductsRequest 
	return request, nil
}

func makeDeleteProductEndpoint(ds DeleteService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(deleteProductsRequest)
		result := ds.deleteProduct(req.ID, req.Q, req.A)
		return deleteProductResponse{result}, nil
	}
}

func decodeDeleteProductRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request deleteProductsRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

