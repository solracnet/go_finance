package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/solracnet/go_finance_backend/db/sqlc"
	"github.com/solracnet/go_finance_backend/util"
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
	err := util.GetAndVerifyToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	var req createAccountRequest
	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var categoryId = req.CategoryID
	var accountType = req.Type
	category, err := server.store.GetCategoryById(ctx, categoryId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}
	if category.Type != accountType {
		ctx.JSON(http.StatusBadRequest, "Category type is not equal to account type")
		return
	}

	arg := db.CreateAccountParams{
		UserID:      req.UserID,
		CategoryID:  categoryId,
		Title:       req.Title,
		Type:        accountType,
		Description: req.Description,
		Value:       req.Value,
		Date:        req.Date,
	}

	user, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, user)
}

type getAccountRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) GetAccount(ctx *gin.Context) {
	err := util.GetAndVerifyToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	var req getAccountRequest
	err = ctx.ShouldBindUri(&req)
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
	err := util.GetAndVerifyToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	var req getAccountRequest
	err = ctx.ShouldBindUri(&req)
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
	err := util.GetAndVerifyToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	var req updateAccountRequest
	err = ctx.ShouldBindJSON(&req)
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
	UserID      int32     `form:"user_id" json:"user_id" binding:"required"`
	Type        string    `form:"type" json:"type" binding:"required"`
	CategoryID  int32     `form:"category_id" json:"category_id"`
	Title       string    `form:"title" json:"title"`
	Description string    `form:"description" json:"description"`
	Date        time.Time `form:"date" json:"date"`
}

func (server *Server) GetAccounts(ctx *gin.Context) {
	err := util.GetAndVerifyToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	var req getAccountsRequest
	err = ctx.ShouldBindQuery(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := db.GetAccountsParams{
		UserID:      req.UserID,
		Type:        req.Type,
		CategoryID:  sql.NullInt32{Int32: req.CategoryID, Valid: req.CategoryID > 0},
		Title:       req.Title,
		Description: req.Description,
		Date:        sql.NullTime{Time: req.Date, Valid: !req.Date.IsZero()},
	}
	accounts, err := server.store.GetAccounts(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, accounts)
}

type getAccountGraphRequest struct {
	UserID int32  `uri:"user_id" binding:"required"`
	Type   string `uri:"type" binding:"required"`
}

func (server *Server) GetAccountGraph(ctx *gin.Context) {
	err := util.GetAndVerifyToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	var req getAccountGraphRequest
	err = ctx.ShouldBindUri(&req)
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
	err := util.GetAndVerifyToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	var req getAccountReportsRequest
	err = ctx.ShouldBindJSON(&req)
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
