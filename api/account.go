package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/solracnet/go_finance_backend/db/sqlc"
)

type createAccountRequest struct {
	UserID      int32     `json:"user_id"`
	CategoryID  int32     `json:"category_id"`
	Title       string    `json:"title"`
	Type        string    `json:"type"`
	Description string    `json:"description"`
	Value       int32     `json:"value"`
	Date        time.Time `json:"date"`
}

func (server *Server) CreateAccount(ctx *gin.Context) {
	var req createAccountRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	arg := db.CreateAccountParams{
		UserID:      req.UserID,
		CategoryID:  req.CategoryID,
		Title:       req.Title,
		Type:        req.Type,
		Description: req.Description,
		Value:       req.Value,
		Date:        req.Date,
	}

	user, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, user)
}

type getAccountRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) GetAccount(ctx *gin.Context) {
	var req getAccountRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	account, err := server.store.GetAccountById(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, err)
		}
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, account)
}

func (server *Server) DeleteAccount(ctx *gin.Context) {
	var req getAccountRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	err = server.store.DeleteAccount(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.Status(http.StatusNoContent)
}

type updateAccountRequest struct {
	ID          int32  `json:"id" binding:"required"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Value       int32  `json:"value"`
}

func (server *Server) UpdateAccount(ctx *gin.Context) {
	var req updateAccountRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	arg := db.UpdateAccountParams{
		ID:          req.ID,
		Title:       req.Title,
		Description: req.Description,
		Value:       req.Value,
	}

	category, err := server.store.UpdateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, category)
}

type getAccountsRequest struct {
	UserID      int32     `json:"user_id"`
	CategoryID  int32     `json:"category_id"`
	Type        string    `json:"type"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Date        time.Time `json:"date"`
}

func (server *Server) GetAccounts(ctx *gin.Context) {
	var req getAccountsRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	arg := db.GetAccountsParams{
		UserID:      req.UserID,
		CategoryID:  req.CategoryID,
		Type:        req.Type,
		Title:       req.Title,
		Description: req.Description,
		Date:        req.Date,
	}

	category, err := server.store.GetAccounts(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, err)
		}
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, category)
}

type getAccountGraphRequest struct {
	UserID int32  `uri:"user_id" binding:"required"`
	Type   string `uri:"type" binding:"required"`
}

func (server *Server) GetAccountGraph(ctx *gin.Context) {
	var req getAccountGraphRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	arg := db.GetAccountsGraphParams{
		UserID: req.UserID,
		Type:   req.Type,
	}

	sumReports, err := server.store.GetAccountsGraph(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, sumReports)
}

type getAccountReportsRequest struct {
	UserID int32  `json:"user_id" binding:"required"`
	Type   string `json:"type" binding:"required"`
}

func (server *Server) GetAccountReports(ctx *gin.Context) {
	var req getAccountReportsRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	arg := db.GetAccountsReportsParams{
		UserID: req.UserID,
		Type:   req.Type,
	}

	countReports, err := server.store.GetAccountsReports(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, countReports)
}
