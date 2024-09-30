package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jasurxaydarov/book_shop_api_getway/api/handlers"
	"github.com/jasurxaydarov/book_shop_api_getway/api/middlewares"
	"github.com/jasurxaydarov/book_shop_api_getway/redis"
	"github.com/jasurxaydarov/book_shop_api_getway/service"

	"github.com/saidamir98/udevs_pkg/logger"
)

type Options struct {
	Service service.ServiceManagerI
	Log     logger.LoggerI
	Cache   redis.RedisRepoI
}

func Api(o Options) *gin.Engine {

	h := handlers.NewHandlers(o.Service, o.Log, o.Cache)

	engine := gin.Default()

	api := engine.Group("/api")
	us := api.Group("/us")

	us.Use(middlewares.AuthMiddlewareUser())
	{
		us.GET("/user/:id", h.GetUserById)
		us.POST("/update", h.UpdateUser)
		us.POST("/delete", h.DeleteUser)

		//order
		us.POST("/order", h.CreateOrder)
		us.GET("/order/:id", h.GetOrderById)
		us.POST("/orders", h.GetOrders)
		us.POST("/order_update", h.UpdateOrder)
		us.POST("/order_delete", h.DeleteOrder)



		//Author
		us.GET("/auth/:id", h.GetAuthById)
		us.POST("/auths", h.GetAuths)


		// book
		us.GET("/book/:id", h.GetBookById)
		us.POST("/books", h.GetBooks)


		// orderItem
		us.POST("/order_item", h.CreateOrderItem)
		us.GET("/order_item/:id", h.GetOrderItemById)
		us.GET("/order_item_id/:id", h.GetOrderItemById)
		us.POST("/order_items", h.GetOrderItems)
		us.POST("/order_item_update", h.UpdateOrderItem)
		us.POST("/order_item_delete", h.DeleteOrderItem)


	}

	adm := api.Group("/adm")

	adm.Use(middlewares.AuthMiddlewareAdmin())
	{
		// author
		adm.POST("/auth", h.CreateAuth)
		adm.GET("/auth/:id", h.GetAuthById)
		adm.POST("/auths", h.GetAuths)
		adm.POST("/auth_update", h.UpdateAuth)
		adm.POST("/auth_delete", h.DeleteAuth)


		/////
		adm.POST("/update", h.UpdateUser)
		adm.POST("/delete", h.DeleteUser)

		////////
		adm.POST("/update_user", h.AdmUpdateUser)
		adm.POST("/delete_user", h.AdmDeleteUser)
		adm.POST("/get_users", h.GetUsers)





		//category
		adm.POST("/category", h.CreateCategory)
		adm.GET("/category/:id", h.GetCategoryById)
		adm.POST("/categories", h.GetCategories)
		adm.POST("/category_update", h.UpdateCategory)
		adm.POST("/category_delete", h.DeleteCategory)


		//book
		adm.POST("/book", h.CreateBook)
		adm.GET("/book/:id", h.GetBookById)
		adm.POST("/books", h.GetBooks)
		adm.POST("/book_update", h.UpdateBook)
		adm.POST("/book_delete", h.DeleteBook)

		///order
		adm.POST("/order", h.CreateOrder)
		adm.GET("/order/:id", h.GetOrderById)
		adm.POST("/orders", h.GetOrders)
		adm.POST("/order_update", h.UpdateOrder)
		adm.POST("/order_delete", h.DeleteOrder)

		/// order items
		adm.POST("/order_item", h.CreateOrderItem)
		adm.GET("/order_item/:id", h.GetOrderItemById)
		adm.GET("/order_item_id/:id", h.GetOrderItemById)
		adm.POST("/order_items", h.GetOrderItems)
		adm.POST("/order_item_update", h.UpdateOrderItem)
		adm.POST("/order_item_delete", h.DeleteOrderItem)

	}

	all := api.Group("/all")

	{
		all.POST("/check-user", h.CheckUser) //completed
		all.POST("/sign-up", h.SignUp)       //completed
		all.POST("/sign-in", h.SigIn)        //completed

	}
	return engine

}
