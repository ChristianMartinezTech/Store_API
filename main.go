// Store API
package main

import (
	"net/http"

	// "net/http"
	// "errors"
	"github.com/gin-gonic/gin"
)

/* Future categories:
Technology sttruct
Books struct
Fitness struct
Beauty struct
Software struct
Toys struct
*/

type Product struct {
	ID              int     `json:"id"`
	ProductName     string  `json:"productname"`
	ProductCategory string  `json:"productcategory"`
	Price           float32 `json:"price"`
	Quantity        int     `json:"quantity"`
	Description     string  `json:"description"`
}

// Inventory (I'll move this later on to a data base with CRUD functonality)
var Products = []Product{
	{ID: 1,
		ProductName:     "MacBook Air 2020",
		ProductCategory: "Technology",
		Price:           999.99,
		Quantity:        8,
		Description:     "",
	}, {ID: 2,
		ProductName:     "MacBook Air 2022",
		ProductCategory: "Technology",
		Price:           1299.99,
		Quantity:        3,
		Description:     "",
	}, {ID: 3,
		ProductName:     "MacBook Pro M1 Pro",
		ProductCategory: "Technology",
		Price:           1499.99,
		Quantity:        6,
		Description:     "",
	},
}

// func that handles the Products endpoint
func getProducts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, Products)
}

func main() {
	router := gin.Default()
	router.GET("/Products", getProducts)
	router.Run("localhost:8080")
}
