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

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
func DecodeUserReq(_ context.Context, r *http.Request) (interface{}, error) {
	var req UserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, err
	}
	return req, nil
}
