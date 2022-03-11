package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/kazuki-sep27/simple_bank_go/db/sqlc"
)

type CreateAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof= USD THB EUR"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req CreateAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest,errorResponse(err))
		return
	}

	arg := db.CreateAccountsParams{
		Owner: req.Owner,
		Balance: 0,
		Currency: req.Currency,
	}

	result, err := server.store.CreateAccounts(ctx, arg)
	
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}

	accountID, err := result.LastInsertId()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}

	account, err := server.store.GetAccountByID(ctx,accountID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK,account)
}