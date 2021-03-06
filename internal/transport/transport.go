package transport

import (
	"context"
	"encoding/json"
	"net/http"

	"discount_service/internal/dto"
	"discount_service/internal/endpoint"
	"discount_service/internal/model"

	kitendpoint "github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHTTPServer(ctx context.Context, endpoints endpoint.Endpoints) http.Handler {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(encodeError),
	}

	r := mux.NewRouter()
	r.Use(commonMiddleware)

	r.Methods("POST").Path("/").Handler(httptransport.NewServer(
		endpoints.ChargeWallet,
		DecodeChargeRequest,
		encodeResponse,
		options...,
	))

	r.Methods("POST").Path("/initiate").Handler(httptransport.NewServer(
		endpoints.InitiateDiscounts,
		DecodeInitiateDiscounts,
		encodeResponse,
		options...,
	))

	r.Methods("GET").Path("/{id}").Handler(httptransport.NewServer(
		endpoints.GetDiscountsByID,
		DecodeGetDiscountsByIDRequest,
		encodeResponse,
		options...,
	))

	r.Methods("GET").Path("/").Handler(httptransport.NewServer(
		endpoints.GetDiscounts,
		DecodeGetDiscountsRequest,
		encodeResponse,
		options...,
	))

	return r
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.WriteHeader(model.ErrToHTTPStatus(err))
	json.NewEncoder(w).Encode(
		dto.Error{
			Message:   err.Error(),
			ErrorCode: model.ErrToCode(err),
		},
	)
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(kitendpoint.Failer); ok && e.Failed() != nil {
		encodeError(ctx, e.Failed(), w)
		return nil
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
