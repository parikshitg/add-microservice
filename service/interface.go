package service

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/parikshitg/add-microservice/api"
)

// AddService interface
type AddService interface {
	Add(context.Context, api.AddRequest) (api.AddResponse, error)
}

// Add - Wrapper for AddService Interface
type Add struct{}

// EncodeResponse -
func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
