package route

import (
	"demo_mogoDB/api"
	"demo_mogoDB/repository"

	"github.com/globalsign/mgo"

	"github.com/gin-gonic/gin"
)

func NewRouteProduct(route *gin.Engine, connectionDB *mgo.Session) {
	productRepository := repository.ProductRepositoryMongo{
		ConnectionDB: connectionDB,
	}
	productAPI := api.ProductAPI{
		ProductRepository: &productRepository,
	}
	route.GET("api/v1/product", productAPI.ProductListHandler)
	route.POST("api/v1/product", productAPI.AddProductHandeler)
	route.PUT("api/v1/product/:product_id", productAPI.EditProducNametHandler)
	route.DELETE("api/v1/product/:product_id", productAPI.DeleteProductByIDHandler)
}
