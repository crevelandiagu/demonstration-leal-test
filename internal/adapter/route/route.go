package route

import (
	_ "demonstration-leal-test/docs"
	"demonstration-leal-test/internal/adapter/handler"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

// Router is a wrapper for HTTP router
type Router struct {
	*gin.Engine
}

func NewRouter(
	lealHandler handler.LealHandler,
) (*Router, error) {
	router := gin.Default()

	// add swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/ping", Ping)

	baseRouter := router.Group("/v1")
	//tagsRouter := baseRouter.Group("/user")
	baseRouter.POST("/user", lealHandler.SaveUser)
	baseRouter.GET("/user/:id", lealHandler.FindOneUser)
	baseRouter.POST("/commerce", lealHandler.SaveCommerce)
	baseRouter.POST("/commerce/:id/branche", lealHandler.SaveBrancheCommerce)
	baseRouter.GET("/commerce/:id", lealHandler.FindCommerce)
	baseRouter.POST("/commerce/:id/purchase-transaction", lealHandler.SavePurchaseTransaction)
	baseRouter.GET("/commerce/:id/purchase-transaction", lealHandler.FindPurchaseTransaction)
	baseRouter.POST("/commerce/:id/branche/:idBranche/campaign", lealHandler.SaveCampaignBranche)
	baseRouter.GET("/commerce/:id/branche/:idBranche/campaign", lealHandler.FindCampaignBranche)
	baseRouter.POST("/commerce/:id/redeem/", lealHandler.SaveRedeem)
	baseRouter.GET("/commerce/:id/redeem/ ", lealHandler.FindRedeem)
	baseRouter.POST("/commerce/:id/campaign", lealHandler.SaveCampaignCommerce)
	baseRouter.GET("/commerce/:id/campaign", lealHandler.FindCampaignCommerce)

	return &Router{
		router,
	}, nil
}

// Serve starts the HTTP server
func (r *Router) Serve(listenAddr string) error {
	return r.Run(listenAddr)
}

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags healthcheck
// @Accept json
// @Produce json
// @Success 200 {string} Ping
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, "pong")
}
