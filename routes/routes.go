package routes

import (
    "shoe_inventory/controllers"
    "github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    router := gin.Default()

    shoeRoutes := router.Group("/shoes")
    {
        shoeRoutes.GET("/", controllers.GetShoes)
        shoeRoutes.POST("/", controllers.CreateShoe)
        shoeRoutes.GET("/:id", controllers.GetShoeByID)
        shoeRoutes.PUT("/:id", controllers.UpdateShoe)
        shoeRoutes.DELETE("/:id", controllers.DeleteShoe)
    }

    return router
}
