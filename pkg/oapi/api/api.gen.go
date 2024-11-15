// Package spec provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package spec

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/oapi-codegen/runtime"
	strictnethttp "github.com/oapi-codegen/runtime/strictmiddleware/nethttp"
)

// Defines values for TransactionState.
const (
	Lose TransactionState = "lose"
	Win  TransactionState = "win"
)

// Defines values for PostUserUserIdTransactionParamsSourceType.
const (
	Game    PostUserUserIdTransactionParamsSourceType = "game"
	Payment PostUserUserIdTransactionParamsSourceType = "payment"
	Server  PostUserUserIdTransactionParamsSourceType = "server"
)

// Error defines model for Error.
type Error struct {
	Message string `json:"message"`
}

// Transaction defines model for Transaction.
type Transaction struct {
	Amount        string           `json:"amount" validate:"required,numeric"`
	State         TransactionState `json:"state" validate:"required,oneof=win lose"`
	TransactionId string           `json:"transactionId" validate:"required,max=256"`
}

// TransactionState defines model for Transaction.State.
type TransactionState string

// UserBalance defines model for UserBalance.
type UserBalance struct {
	Balance string `json:"balance"`
	UserId  uint64 `json:"userId"`
}

// PostUserUserIdTransactionParams defines parameters for PostUserUserIdTransaction.
type PostUserUserIdTransactionParams struct {
	SourceType PostUserUserIdTransactionParamsSourceType `json:"Source-Type" validate:"required,oneof=game server payment"`
}

// PostUserUserIdTransactionParamsSourceType defines parameters for PostUserUserIdTransaction.
type PostUserUserIdTransactionParamsSourceType string

// PostUserUserIdTransactionJSONRequestBody defines body for PostUserUserIdTransaction for application/json ContentType.
type PostUserUserIdTransactionJSONRequestBody = Transaction

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Get user balance
	// (GET /user/{userId}/balance)
	GetUserUserIdBalance(w http.ResponseWriter, r *http.Request, userId uint64)
	// Update user balance
	// (POST /user/{userId}/transaction)
	PostUserUserIdTransaction(w http.ResponseWriter, r *http.Request, userId uint64, params PostUserUserIdTransactionParams)
}

// Unimplemented server implementation that returns http.StatusNotImplemented for each endpoint.

type Unimplemented struct{}

// Get user balance
// (GET /user/{userId}/balance)
func (_ Unimplemented) GetUserUserIdBalance(w http.ResponseWriter, r *http.Request, userId uint64) {
	w.WriteHeader(http.StatusNotImplemented)
}

// Update user balance
// (POST /user/{userId}/transaction)
func (_ Unimplemented) PostUserUserIdTransaction(w http.ResponseWriter, r *http.Request, userId uint64, params PostUserUserIdTransactionParams) {
	w.WriteHeader(http.StatusNotImplemented)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandlerFunc   func(w http.ResponseWriter, r *http.Request, err error)
}

type MiddlewareFunc func(http.Handler) http.Handler

// GetUserUserIdBalance operation middleware
func (siw *ServerInterfaceWrapper) GetUserUserIdBalance(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "userId" -------------
	var userId uint64

	err = runtime.BindStyledParameterWithOptions("simple", "userId", chi.URLParam(r, "userId"), &userId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "userId", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetUserUserIdBalance(w, r, userId)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

// PostUserUserIdTransaction operation middleware
func (siw *ServerInterfaceWrapper) PostUserUserIdTransaction(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var err error

	// ------------- Path parameter "userId" -------------
	var userId uint64

	err = runtime.BindStyledParameterWithOptions("simple", "userId", chi.URLParam(r, "userId"), &userId, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationPath, Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "userId", Err: err})
		return
	}

	// Parameter object where we will unmarshal all parameters from the context
	var params PostUserUserIdTransactionParams

	headers := r.Header

	// ------------- Required header parameter "Source-Type" -------------
	if valueList, found := headers[http.CanonicalHeaderKey("Source-Type")]; found {
		var SourceType PostUserUserIdTransactionParamsSourceType
		n := len(valueList)
		if n != 1 {
			siw.ErrorHandlerFunc(w, r, &TooManyValuesForParamError{ParamName: "Source-Type", Count: n})
			return
		}

		err = runtime.BindStyledParameterWithOptions("simple", "Source-Type", valueList[0], &SourceType, runtime.BindStyledParameterOptions{ParamLocation: runtime.ParamLocationHeader, Explode: false, Required: true})
		if err != nil {
			siw.ErrorHandlerFunc(w, r, &InvalidParamFormatError{ParamName: "Source-Type", Err: err})
			return
		}

		params.SourceType = SourceType

	} else {
		err := fmt.Errorf("Header parameter Source-Type is required, but not found")
		siw.ErrorHandlerFunc(w, r, &RequiredHeaderError{ParamName: "Source-Type", Err: err})
		return
	}

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.PostUserUserIdTransaction(w, r, userId, params)
	}))

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler.ServeHTTP(w, r.WithContext(ctx))
}

type UnescapedCookieParamError struct {
	ParamName string
	Err       error
}

func (e *UnescapedCookieParamError) Error() string {
	return fmt.Sprintf("error unescaping cookie parameter '%s'", e.ParamName)
}

func (e *UnescapedCookieParamError) Unwrap() error {
	return e.Err
}

type UnmarshalingParamError struct {
	ParamName string
	Err       error
}

func (e *UnmarshalingParamError) Error() string {
	return fmt.Sprintf("Error unmarshaling parameter %s as JSON: %s", e.ParamName, e.Err.Error())
}

func (e *UnmarshalingParamError) Unwrap() error {
	return e.Err
}

type RequiredParamError struct {
	ParamName string
}

func (e *RequiredParamError) Error() string {
	return fmt.Sprintf("Query argument %s is required, but not found", e.ParamName)
}

type RequiredHeaderError struct {
	ParamName string
	Err       error
}

func (e *RequiredHeaderError) Error() string {
	return fmt.Sprintf("Header parameter %s is required, but not found", e.ParamName)
}

func (e *RequiredHeaderError) Unwrap() error {
	return e.Err
}

type InvalidParamFormatError struct {
	ParamName string
	Err       error
}

func (e *InvalidParamFormatError) Error() string {
	return fmt.Sprintf("Invalid format for parameter %s: %s", e.ParamName, e.Err.Error())
}

func (e *InvalidParamFormatError) Unwrap() error {
	return e.Err
}

type TooManyValuesForParamError struct {
	ParamName string
	Count     int
}

func (e *TooManyValuesForParamError) Error() string {
	return fmt.Sprintf("Expected one value for %s, got %d", e.ParamName, e.Count)
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL          string
	BaseRouter       chi.Router
	Middlewares      []MiddlewareFunc
	ErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	if options.ErrorHandlerFunc == nil {
		options.ErrorHandlerFunc = func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandlerFunc:   options.ErrorHandlerFunc,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/user/{userId}/balance", wrapper.GetUserUserIdBalance)
	})
	r.Group(func(r chi.Router) {
		r.Post(options.BaseURL+"/user/{userId}/transaction", wrapper.PostUserUserIdTransaction)
	})

	return r
}

type GetUserUserIdBalanceRequestObject struct {
	UserId uint64 `json:"userId"`
}

type GetUserUserIdBalanceResponseObject interface {
	VisitGetUserUserIdBalanceResponse(w http.ResponseWriter) error
}

type GetUserUserIdBalance200JSONResponse UserBalance

func (response GetUserUserIdBalance200JSONResponse) VisitGetUserUserIdBalanceResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type GetUserUserIdBalancedefaultJSONResponse struct {
	Body       Error
	StatusCode int
}

func (response GetUserUserIdBalancedefaultJSONResponse) VisitGetUserUserIdBalanceResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)

	return json.NewEncoder(w).Encode(response.Body)
}

type PostUserUserIdTransactionRequestObject struct {
	UserId uint64 `json:"userId"`
	Params PostUserUserIdTransactionParams
	Body   *PostUserUserIdTransactionJSONRequestBody
}

type PostUserUserIdTransactionResponseObject interface {
	VisitPostUserUserIdTransactionResponse(w http.ResponseWriter) error
}

type PostUserUserIdTransaction200JSONResponse UserBalance

func (response PostUserUserIdTransaction200JSONResponse) VisitPostUserUserIdTransactionResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	return json.NewEncoder(w).Encode(response)
}

type PostUserUserIdTransactiondefaultJSONResponse struct {
	Body       Error
	StatusCode int
}

func (response PostUserUserIdTransactiondefaultJSONResponse) VisitPostUserUserIdTransactionResponse(w http.ResponseWriter) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)

	return json.NewEncoder(w).Encode(response.Body)
}

// StrictServerInterface represents all server handlers.
type StrictServerInterface interface {
	// Get user balance
	// (GET /user/{userId}/balance)
	GetUserUserIdBalance(ctx context.Context, request GetUserUserIdBalanceRequestObject) (GetUserUserIdBalanceResponseObject, error)
	// Update user balance
	// (POST /user/{userId}/transaction)
	PostUserUserIdTransaction(ctx context.Context, request PostUserUserIdTransactionRequestObject) (PostUserUserIdTransactionResponseObject, error)
}

type StrictHandlerFunc = strictnethttp.StrictHTTPHandlerFunc
type StrictMiddlewareFunc = strictnethttp.StrictHTTPMiddlewareFunc

type StrictHTTPServerOptions struct {
	RequestErrorHandlerFunc  func(w http.ResponseWriter, r *http.Request, err error)
	ResponseErrorHandlerFunc func(w http.ResponseWriter, r *http.Request, err error)
}

func NewStrictHandler(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: StrictHTTPServerOptions{
		RequestErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusBadRequest)
		},
		ResponseErrorHandlerFunc: func(w http.ResponseWriter, r *http.Request, err error) {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		},
	}}
}

func NewStrictHandlerWithOptions(ssi StrictServerInterface, middlewares []StrictMiddlewareFunc, options StrictHTTPServerOptions) ServerInterface {
	return &strictHandler{ssi: ssi, middlewares: middlewares, options: options}
}

type strictHandler struct {
	ssi         StrictServerInterface
	middlewares []StrictMiddlewareFunc
	options     StrictHTTPServerOptions
}

// GetUserUserIdBalance operation middleware
func (sh *strictHandler) GetUserUserIdBalance(w http.ResponseWriter, r *http.Request, userId uint64) {
	var request GetUserUserIdBalanceRequestObject

	request.UserId = userId

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.GetUserUserIdBalance(ctx, request.(GetUserUserIdBalanceRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "GetUserUserIdBalance")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(GetUserUserIdBalanceResponseObject); ok {
		if err := validResponse.VisitGetUserUserIdBalanceResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}

// PostUserUserIdTransaction operation middleware
func (sh *strictHandler) PostUserUserIdTransaction(w http.ResponseWriter, r *http.Request, userId uint64, params PostUserUserIdTransactionParams) {
	var request PostUserUserIdTransactionRequestObject

	request.UserId = userId
	request.Params = params

	var body PostUserUserIdTransactionJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		sh.options.RequestErrorHandlerFunc(w, r, fmt.Errorf("can't decode JSON body: %w", err))
		return
	}
	request.Body = &body

	handler := func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (interface{}, error) {
		return sh.ssi.PostUserUserIdTransaction(ctx, request.(PostUserUserIdTransactionRequestObject))
	}
	for _, middleware := range sh.middlewares {
		handler = middleware(handler, "PostUserUserIdTransaction")
	}

	response, err := handler(r.Context(), w, r, request)

	if err != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, err)
	} else if validResponse, ok := response.(PostUserUserIdTransactionResponseObject); ok {
		if err := validResponse.VisitPostUserUserIdTransactionResponse(w); err != nil {
			sh.options.ResponseErrorHandlerFunc(w, r, err)
		}
	} else if response != nil {
		sh.options.ResponseErrorHandlerFunc(w, r, fmt.Errorf("unexpected response type: %T", response))
	}
}
