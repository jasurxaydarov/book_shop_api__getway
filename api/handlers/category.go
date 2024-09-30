package handlers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jasurxaydarov/book_shop_api_getway/genproto/product_service"
)

func (h *Handler) CreateCategory(ctx *gin.Context) {
	var req product_service.CategoryCreateReq

	ctx.BindJSON(&req)

	resp, err := h.service.GetProductSevice().CreateCategory(context.Background(), &req)

	if err != nil {
		h.log.Error(err.Error())
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(201, resp)

}

func (h *Handler) GetCategoryById(ctx *gin.Context) {

	var req product_service.GetByIdReq

	req.Id = ctx.Param("id")

	resp, err := h.service.GetProductSevice().GetCategory(context.Background(), &req)

	if err != nil {
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(201, resp)	
}




func (h *Handler) GetCategories(ctx *gin.Context) {

	var req product_service.GetListReq

	ctx.BindJSON(&req)

	resp, err := h.service.GetProductSevice().GetCategories(context.Background(), &req)

	if err != nil {
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(201, resp)

}

func (h *Handler) UpdateCategory(ctx *gin.Context) {

	req := product_service.CategoryUpdateReq{}


	ctx.BindJSON(&req)


	res, err := h.service.GetProductSevice().UpdateCategory(context.Background(), &req)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201,res)

}

func (h *Handler) DeleteCategory(ctx *gin.Context) {

	req := product_service.DeleteReq{}


	ctx.BindJSON(&req)

	_, err := h.service.GetProductSevice().DeleteCategory(context.Background(), &req)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201,"succssfully deleted")

}