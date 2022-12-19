package routes

import (
	"recipe-api/handlers"

	"github.com/gin-gonic/gin"
)

func NewRoutes() *gin.Engine {
	router := gin.Default()
	// レシピ作成
	router.POST("/recipes", handlers.NewRecipeHandler)
	router.GET("/recipes", handlers.ListRecipesHandler)
	router.PUT("/recipes/:id", handlers.UpdateRecipeHandler)
	router.DELETE("/recipes/:id", handlers.DeleteRecipeHandler)
	router.GET("/recipes/search", handlers.SearchRecipesHandler)
	return router
}
