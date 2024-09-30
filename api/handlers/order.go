package handlers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jasurxaydarov/book_shop_api_getway/genproto/product_service"
)

func (h *Handler) CreateOrder(ctx *gin.Context) {
	var req product_service.OrderCreateReq

	ctx.BindJSON(&req)

	resp, err := h.service.GetProductSevice().CreateOrder(context.Background(), &req)

	if err != nil {
		h.log.Error(err.Error())
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(201, resp)

}

func (h *Handler) GetOrderById(ctx *gin.Context) {

	var req product_service.GetByIdReq

	req.Id = ctx.Param("id")

	resp, err := h.service.GetProductSevice().GetOrder(context.Background(), &req)

	if err != nil {
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(201, resp)

	
}



func (h *Handler) GetOrders(ctx *gin.Context) {

	var req product_service.GetListReq

	ctx.BindJSON(&req)

	resp, err := h.service.GetProductSevice().GetOrders(context.Background(), &req)

	if err != nil {
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(201, resp)

}

func (h *Handler) UpdateOrder(ctx *gin.Context) {

	req := product_service.OrderUpdateReq{}


	ctx.BindJSON(&req)


	res, err := h.service.GetProductSevice().Updateorder(context.Background(), &req)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201,res)

}

func (h *Handler) DeleteOrder(ctx *gin.Context) {

	req := product_service.DeleteReq{}


	ctx.BindJSON(&req)

	_, err := h.service.GetProductSevice().DeleteOrder(context.Background(), &req)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201,"succssfully deleted")

}