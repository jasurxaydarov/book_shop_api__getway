package handlers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jasurxaydarov/book_shop_api_getway/genproto/product_service"
	"github.com/saidamir98/udevs_pkg/logger"
)

func (h *Handler) CreateBook(ctx *gin.Context) {

	var req product_service.BookCreateReq

	ctx.BindJSON(&req)

	resp, err := h.service.GetProductSevice().CreateBook(context.Background(), &req)

	if err != nil {
		h.log.Error("err on  GetBookService().CreateBook", logger.Error(err))
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(201, resp)
}

func (h *Handler) GetBookById(ctx *gin.Context) {

	var req product_service.GetByIdReq

	req.Id = ctx.Param("id")

	resp, err := h.service.GetProductSevice().GetBook(context.Background(), &req)

	if err != nil {
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(201, resp)

}





func (h *Handler) GetBooks(ctx *gin.Context) {

	var req product_service.GetListReq

	ctx.BindJSON(&req)

	resp, err := h.service.GetProductSevice().GetBooks(context.Background(), &req)

	if err != nil {
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(201, resp)

}

func (h *Handler) UpdateBook(ctx *gin.Context) {

	req := product_service.BookUpdateReq{}


	ctx.BindJSON(&req)


	res, err := h.service.GetProductSevice().UpdateBook(context.Background(), &req)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201,res)

}

func (h *Handler) DeleteBook(ctx *gin.Context) {

	req := product_service.DeleteReq{}


	ctx.BindJSON(&req)

	_, err := h.service.GetProductSevice().DeleteBook(context.Background(), &req)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201,"succssfully deleted")

}