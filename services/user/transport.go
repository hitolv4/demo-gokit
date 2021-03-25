package user

import (
	"context"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/memeoAmazonas/demo-2/common"
	"net/http"
)

func NewHttpServer(_ context.Context, endpoint Endpoint) http.Handler {

	r := mux.NewRouter()
	userRoute := r.PathPrefix("/users").Subrouter()
	userRoute.Use(common.CommonMiddleware)
	userRoute.Methods("POST").Path("/new").Handler(httptransport.NewServer(
		endpoint.CreateUser,
		DecodeUserReq,
		EncodeResponse,
	))
	userRoute.Methods("GET").Path("/").Handler(httptransport.NewServer(
		endpoint.GetAllUser,
		DecodeAll,
		EncodeResponse,
	))
	userRoute.Methods("GET").Path("/{id}").Handler(httptransport.NewServer(
		endpoint.GetUserById,
		DecodeId,
		EncodeResponse,
	))
	userRoute.Methods("DELETE").Path("/{id}").Handler(httptransport.NewServer(
		endpoint.Delete,
		DecodeId,
		EncodeResponse,
	))
	return userRoute
}
