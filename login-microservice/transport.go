package main 

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
)
/*
	In this middleware layer I define the different types of communications and structure of the 
	messages that can be made to and from the RPC
*/
type loginRequest struct {
	U string `json: "u"`
	P string `json: "p"`
	A bool `json: "a"`
}

type loginResponse struct {
	R bool `json:"r"`
}
func makeLoginEndpoint(ls LoginService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(loginRequest)
		result := ls.validateCredentials(req.U, req.P, req.A)
		return loginResponse{result}, nil
	}
}

func decodeLoginRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request loginRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
