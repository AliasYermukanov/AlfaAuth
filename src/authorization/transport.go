package authorization

import (
	"context"
	"encoding/json"
	"github.com/AliasYermukanov/AlfaAuth/src/errors"
	"github.com/AliasYermukanov/AlfaAuth/src/middleware"
	kitlog "github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

func MakeHandler(ss Service, logger kitlog.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(encodeError),
	}

	r := mux.NewRouter()

	createUser := kithttp.NewServer(
		createUserEndpoint(ss),
		decodeCreateUserRequest,
		encodeResponse,
		opts...,
	)

	authentication := kithttp.NewServer(
		authenticationEndpoint(ss),
		decodeCreateUserRequest,
		encodeResponse,
		opts...,
	)

	updateUserData := kithttp.NewServer(
		UpdateUserEndpoint(ss),
		decodeUpdateUserRequest,
		encodeResponse,
		opts...,
	)

	r.Handle("/v1/auth/create-user", createUser).Methods("POST")
	r.Handle("/v1/auth/authentication", authentication).Methods("POST")
	r.Handle("/v1/auth/update-userdata",middleware.AuthMiddleware(updateUserData, "SUPER_SU", "UPDATE")).Methods("POST")

	return r
}

func decodeUpdateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil,nil
}

func decodeCreateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var body createUserRequest

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		errors.InvalidCharacter.DeveloperMessage = err.Error()
		return nil, errors.InvalidCharacter
	}


	return body, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

type errorer interface {
	error() error
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	switch err {
	case errors.InvalidPhoneNumber:
		w.WriteHeader(http.StatusUnprocessableEntity)
	case errors.InvalidCharacter:
		w.WriteHeader(http.StatusBadRequest)
	case errors.AccessDenied:
		w.WriteHeader(http.StatusForbidden)
	case errors.NoFound:
		w.WriteHeader(http.StatusNotFound)
	case errors.ElasticConnectError:
		w.WriteHeader(http.StatusServiceUnavailable)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json; charset= utf-8")
	json.NewEncoder(w).Encode(err)
}
