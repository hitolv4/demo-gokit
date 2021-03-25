package user

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
)

type UserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type IdRequest struct {
	Id string `json:"id"`
}

type UserResponse struct {
	Message  string `json:"message,omitempty"`
	UserList []User `json:"user_list,omitempty"`
	User     *User  `json:"user,omitempty"`
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

func DecodeId(_ context.Context, r *http.Request) (interface{}, error) {
	var req IdRequest
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		return nil, errors.New("bad request id is required")
	}
	req.Id = id
	return req, nil
}
func DecodeAll(_ context.Context, _ *http.Request) (interface{}, error) {
	return nil, nil
}
