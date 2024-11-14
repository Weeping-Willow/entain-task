package api

import (
	"context"
	"log/slog"
	"strconv"

	spec "github.com/Weeping-Willow/entain-task/pkg/oapi/api"
)

func (a *Server) GetUserUserIdBalance(ctx context.Context, request spec.GetUserUserIdBalanceRequestObject) (spec.GetUserUserIdBalanceResponseObject, error) {
	userBalance, err := a.balanceService.GetUserBalance(ctx, request.UserId)
	if err != nil {
		return nil, err
	}

	return spec.GetUserUserIdBalance200JSONResponse{
		UserId:  request.UserId,
		Balance: strconv.FormatFloat(userBalance, 'f', 2, 64),
	}, nil
}

func (a *Server) PostUserUserIdTransaction(ctx context.Context, request spec.PostUserUserIdTransactionRequestObject) (spec.PostUserUserIdTransactionResponseObject, error) {
	if err := a.validator.Struct(request); err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	newBalance, err := a.balanceService.PostNewTransaction(ctx, request)
	if err != nil {
		return nil, err
	}

	return spec.PostUserUserIdTransaction200JSONResponse{
		UserId:  request.UserId,
		Balance: strconv.FormatFloat(newBalance, 'f', 2, 64),
	}, nil
}
