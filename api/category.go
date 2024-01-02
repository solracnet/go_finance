package api

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/solracnet/go_finance_backend/db/sqlc"
)

type createCategoryRequest struct {
	UserID      int32  `json:"user_id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Type        string `json:"type" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (server *Server) CreateCategory(ctx *gin.Context) {
	var req createCategoryRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	arg := db.CreateCategoryParams{
		UserID:      req.UserID,
		Title:       req.Title,
		Type:        req.Type,
		Description: req.Description,
	}

	user, err := server.store.CreateCategory(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, user)
}

type getCategoryRequest struct {
	ID int32 `uri:"id" binding:"required"`
}

func (server *Server) GetCategory(ctx *gin.Context) {
	var req getCategoryRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	category, err := server.store.GetCategoryById(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, err)
		}
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, category)
}

func (server *Server) DeleteCategory(ctx *gin.Context) {
	var req getCategoryRequest
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	err = server.store.DeleteCategory(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.Status(http.StatusNoContent)
}

type updateCategoryRequest struct {
	ID          int32  `json:"id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func (server *Server) UpdateCategory(ctx *gin.Context) {
	var req updateCategoryRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	arg := db.UpdateCategoryParams{
		ID:          req.ID,
		Title:       req.Title,
		Description: req.Description,
	}

	category, err := server.store.UpdateCategory(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, category)
}

type getCategoriesRequest struct {
	UserID      int32  `json:"user_id" binding:"required"`
	Type        string `json:"type" binding:"required"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (server *Server) GetCategories(ctx *gin.Context) {
	var req getCategoriesRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}
	var categories []db.Category
	if len(req.Description) == 0 && len(req.Title) == 0 {
		arg := db.GetCategoriesByUserIdAndTypeParams{
			UserID: req.UserID,
			Type:   req.Type,
		}
		categoriesByUserIdAndType, err := server.store.GetCategoriesByUserIdAndType(ctx, arg)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, err)
				return
			}
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		categories = categoriesByUserIdAndType
	}

	if len(req.Description) > 0 && len(req.Title) == 0 {
		arg := db.GetCategoriesByUserIdAndTypeAndDescriptionParams{
			UserID:      req.UserID,
			Type:        req.Type,
			Description: req.Description,
		}
		GetCategoriesByUserIdAndTypeAndDescription, err := server.store.GetCategoriesByUserIdAndTypeAndDescription(ctx, arg)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, err)
				return
			}
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		categories = GetCategoriesByUserIdAndTypeAndDescription
	}

	if len(req.Description) == 0 && len(req.Title) > 0 {
		arg := db.GetCategoriesByUserIdAndTypeAndTitleParams{
			UserID: req.UserID,
			Type:   req.Type,
			Title:  req.Title,
		}
		GetCategoriesByUserIdAndTypeAndTitle, err := server.store.GetCategoriesByUserIdAndTypeAndTitle(ctx, arg)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, err)
				return
			}
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		categories = GetCategoriesByUserIdAndTypeAndTitle
	}

	if len(req.Description) > 0 && len(req.Title) > 0 {
		arg := db.GetCategoriesParams{
			UserID:      req.UserID,
			Type:        req.Type,
			Title:       req.Title,
			Description: req.Description,
		}
		categoryWithAllParams, err := server.store.GetCategories(ctx, arg)
		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusNotFound, err)
				return
			}
			ctx.JSON(http.StatusInternalServerError, err)
			return
		}
		categories = categoryWithAllParams
	}
	ctx.JSON(http.StatusOK, categories)
}
