package user

import (
	"context"
	"encoding/json"
	"net/http"
)

type UserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
type UserResponse struct {
	OK string `json:"ok"`
}

func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
func DecodeUserReq(ctx context.Context, r *http.Request) (interface{}, error) {
	var req UserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}

