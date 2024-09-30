package handlers

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jasurxaydarov/book_shop_api_getway/genproto/product_service"
)

func (h *Handler) CreateOrderItem(ctx *gin.Context) {
	var req product_service.OrderItemCreateReq

	ctx.BindJSON(&req)

	resp, err := h.service.GetProductSevice().CreateOrdered_Item(context.Background(), &req)

	if err != nil {
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(201, resp)

}

func (h *Handler) GetOrderItemById(ctx *gin.Context) {

	var req product_service.GetByIdReq

	req.Id = ctx.Param("id")

	resp, err := h.service.GetProductSevice().GetOrdered_Item(context.Background(), &req)

	if err != nil {
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(201, resp)

}

func (h *Handler) GetOrderItemByOrderId(ctx *gin.Context) {

	var req product_service.GetByIdReq

	req.Id = ctx.Param("id")

	resp, err := h.service.GetProductSevice().GetOrdered_ItemByOrderId(context.Background(), &req)

	if err != nil {
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(201, resp)

}



func (h *Handler) GetOrderItems(ctx *gin.Context) {

	var req product_service.GetListReq

	ctx.BindJSON(&req)

	resp, err := h.service.GetProductSevice().GetOrdered_Items(context.Background(), &req)

	if err != nil {
		ctx.JSON(500, err)
		return
	}

	ctx.JSON(201, resp)

}

func (h *Handler) UpdateOrderItem(ctx *gin.Context) {

	req := product_service.OrderItemUpdate{}


	ctx.BindJSON(&req)


	res, err := h.service.GetProductSevice().UpdateOrdered_Item(context.Background(), &req)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201,res)

}

func (h *Handler) DeleteOrderItem(ctx *gin.Context) {

	req := product_service.DeleteReq{}


	ctx.BindJSON(&req)

	_, err := h.service.GetProductSevice().DeleteOrdered_Item(context.Background(), &req)

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(201,"succssfully deleted")

}