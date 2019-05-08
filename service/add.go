package service

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/parikshitg/add-microservice/api"
	"github.com/urantiatech/kit/endpoint"
)

// Add functionality
func (a Add) Add(ctx context.Context, req api.AddRequest) (api.AddResponse, error) {
	var resp api.AddResponse

	resp.Result = req.X + req.Y

	return resp, nil
}

// MakeAddEndpoint -
func MakeAddEndpoint(svc Add) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(api.AddRequest)
		return svc.Add(ctx, req)
	}
}

// DecodeAddRequest -
func DecodeAddRequest(ctx context.Context, r *http.Request) (interface{}, error) {
	var request api.AddRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}
