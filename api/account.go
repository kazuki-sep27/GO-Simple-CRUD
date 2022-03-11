package api

import (
	"database/sql"
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

type GetAccountRequest struct {
	ID    int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getAccount(ctx *gin.Context) {
	var req GetAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest,errorResponse(err))
		return
	}

	account, err := server.store.GetAccountByID(ctx,req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound,errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK,account)
}

type ListAccountRequest struct {
	PageID int32 `from:"page_id" binding:"required,min=1"`
	PageSize int32 `from:"page_size" binding:"required,min=5,max=20"`
}

func (server *Server) listAccounts(ctx *gin.Context) {
	var req ListAccountRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest,errorResponse(err))
		return
	}

	arg := db.ListAccountsParams{
		Limit: req.PageSize,
		Offset: (req.PageID-1)*req.PageSize,
	}

	accounts, err := server.store.ListAccounts(ctx,arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound,errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError,errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK,accounts)
}