package usecase

import (
	"context"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	r := mux.NewRouter()
	r.Use(commonMiddleware)

	r.Methods("POST").Path("/").Handler(httptransport.NewServer(
		endpoints.ChargeWallet,
		decodeChargeRequest,
		encodeResponse,
	))

	r.Methods("POST").Path("/initiate").Handler(httptransport.NewServer(
		endpoints.InitiateDiscounts,
		decodeInitiateDiscounts,
		encodeResponse,
	))

	r.Methods("GET").Path("/{id}").Handler(httptransport.NewServer(
		endpoints.GetDiscountsByID,
		decodeGetDiscountsByIDRequest,
		encodeResponse,
	))

	r.Methods("GET").Path("/").Handler(httptransport.NewServer(
		endpoints.GetDiscounts,
		decodeGetDiscountsRequest,
		encodeResponse,
	))

	return r
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
