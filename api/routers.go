/*
 * @Author: javohir-a abdusamatovjavohir@gmail.com
 * @Date: 2024-12-14 03:02:33
 * @LastEditors: javohir-a abdusamatovjavohir@gmail.com
 * @LastEditTime: 2024-12-15 06:01:49
 * @FilePath: /lesson3/app/routers.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/udevs/lesson3/api/docs"
	"github.com/udevs/lesson3/api/handlers"
	"github.com/udevs/lesson3/config"
	"go.uber.org/zap"

	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type HttpService struct {
	ordersHandler  *handlers.OrdersHandler
	productHandler *handlers.ProductsHandler
	logger         *zap.Logger
	cfg            *config.Config
}

func NewHttpService(o *handlers.OrdersHandler, p *handlers.ProductsHandler, l *zap.Logger, c *config.Config) *HttpService {
	return &HttpService{
		ordersHandler:  o,
		productHandler: p,
		logger:         l,
		cfg:            c,
	}
}

// @title Product and Orders
// @version 1.0
// @description test
// @BasePath        /
// @schemes         http
// @in              header
func (h *HttpService) Run() error {
	router := gin.Default()

	router.GET("swagger/*any", ginSwagger.WrapHandler(files.Handler))

	product := router.Group("/products")
	{
		product.POST("", h.productHandler.CreateProduct)
		product.GET("", h.productHandler.GetAllProducts)
		product.GET(":id", h.productHandler.GetProductByID)
		product.PUT(":id", h.productHandler.UpdateProduct)
		product.DELETE(":id", h.productHandler.DeleteProduct)
	}

	orders := router.Group("/orders")
	{
		orders.POST("", h.ordersHandler.CreateOrder)
		orders.GET("", h.ordersHandler.GetAllOrders)
		orders.GET(":id", h.ordersHandler.GetOrderByID)
		orders.PUT(":id", h.ordersHandler.UpdateOrder)
		orders.DELETE(":id", h.ordersHandler.DeleteOrder)
		orders.GET("/report", h.ordersHandler.GenerateReport)
	}

	h.logger.Info("Starting server", zap.String("address", h.cfg.Server.Host))
	err := http.ListenAndServe(h.cfg.Server.Host+":"+h.cfg.Server.Port, router)
	if err != nil {
		return err
	}

	return nil
}
