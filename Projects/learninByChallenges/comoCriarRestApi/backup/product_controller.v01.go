package controller

import (
	"go-api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type productController struct {
	//Usecase
}

func NewProductController() productController {
	return productController{}
}

func (p *productController) GetProducts(ctx *gin.Context) {

	products := []model.Product{
		{
			ID:    1,
			Name:  "Batata frita",
			Price: 20,
		},
	}

	ctx.JSON(http.StatusOK, products)
}
