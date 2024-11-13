package api

import (
	"context"
	"log/slog"
	"net/http"

	spec "github.com/Weeping-Willow/entain-task/pkg/oapi/api"
)

func (a *Server) GetUserUserIdBalance(ctx context.Context, request spec.GetUserUserIdBalanceRequestObject) (spec.GetUserUserIdBalanceResponseObject, error) {
	if request.UserId < 0 {
		return spec.GetUserUserIdBalancedefaultJSONResponse{
			Body: spec.Error{
				Message: "Invalid user id",
			},
			StatusCode: http.StatusBadRequest,
		}, nil
	}

	return spec.GetUserUserIdBalancedefaultJSONResponse{
		Body: spec.Error{
			Message: "Not implemented",
		},
		StatusCode: http.StatusInternalServerError,
	}, nil
}

func (a *Server) PostUserUserIdTransaction(ctx context.Context, request spec.PostUserUserIdTransactionRequestObject) (spec.PostUserUserIdTransactionResponseObject, error) {
	if request.UserId < 0 {
		return spec.PostUserUserIdTransactiondefaultJSONResponse{
			Body: spec.Error{
				Message: "Invalid user id",
			},
			StatusCode: http.StatusBadRequest,
		}, nil
	}

	if err := a.validator.Struct(request.Body); err != nil {
		slog.Error(err.Error())
		return spec.PostUserUserIdTransactiondefaultJSONResponse{
			Body: spec.Error{
				Message: err.Error(),
			},
			StatusCode: http.StatusBadRequest,
		}, nil
	}

	return spec.PostUserUserIdTransactiondefaultJSONResponse{
		Body: spec.Error{
			Message: "Not implemented",
		},
		StatusCode: http.StatusInternalServerError,
	}, nil
}
