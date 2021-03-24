package user

import (
	"context"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/memeoAmazonas/demo-2/common"
	"net/http"
)

func NewHttpServer(ctx context.Context, endpoint Endpoint) http.Handler {

	r := mux.NewRouter()
	r.Use(common.CommonMiddleware)
	userRoute := r.PathPrefix("/users").Subrouter()
	userRoute.Methods("POST").Path("/new").Handler(httptransport.NewServer(
		endpoint.CreateUser,
		DecodeUserReq,
		EncodeResponse,
	))
	return userRoute
}
