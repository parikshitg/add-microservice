package api

// AddRequest structure
type AddRequest struct {
	X int `json:"x"`
	Y int `json:"y"`
}

// AddResponse returns x+y
type AddResponse struct {
	Result int `json:"result"`
}
