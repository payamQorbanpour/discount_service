package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"discount_service/internal/endpoint"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHTTPServer(ctx context.Context, endpoints endpoint.Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	r.Methods("POST").Path("/").Handler(httptransport.NewServer(
		endpoints.ChargeWallet,
		DecodeChargeRequest,
		EncodeResponse,
	))

	r.Methods("POST").Path("/initiate").Handler(httptransport.NewServer(
		endpoints.InitiateDiscounts,
		DecodeInitiateDiscounts,
		EncodeResponse,
	))

	r.Methods("GET").Path("/{id}").Handler(httptransport.NewServer(
		endpoints.GetDiscountsByID,
		DecodeGetDiscountsByIDRequest,
		EncodeResponse,
	))

	r.Methods("GET").Path("/").Handler(httptransport.NewServer(
		endpoints.GetDiscounts,
		DecodeGetDiscountsRequest,
		EncodeResponse,
	))

	return r
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
