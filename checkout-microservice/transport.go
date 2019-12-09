package main

import (
	"context"
	"encoding/json"
	"net/http"
	"github.com/go-kit/kit/endpoint"
)

type checkoutRequest struct {
	CID int `json:"customer_id"`
	ID int `json:"product_id"`
	Q int `json:"quantity"`
}

type checkoutResponse struct {
	S bool `json:success`
}

func makeCheckoutEndpoint(cs CheckoutService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(checkoutRequest)
		response := cs.checkout(req.CID, req.ID, req.Q)
		return checkoutResponse{response}, nil
	}
}

func decodeCheckoutRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request checkoutRequest 
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

