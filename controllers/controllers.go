package controllers

import (
	"log"
	"net/http"

	"github.com/19abhishek/e-commerce/models"
	"github.com/gin-gonic/gin"
)

func HashPassword (password string) string {

}

func VerifyPassword (userPassword string, givenPassword string) (bool, string) {}

func SignUp() gin.HandlerFunc {

	return func(c *gin.Context){
		var ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var cancel models.User
		if err  := c.BindJSON(&user); err != nil {
			c.JSON{http.StatusBadRequest, gin.H("error": err.Error())}
		}

		validationErr := Validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr})
		}

		count, err != UserCollection.CountDocuments(ctx, bson.M{"error": user.Email})
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error"})
		}
	}
}

func Login() gin.HandlerFunc {}

func ProductViewAdmin() gin.HandlerFunc {}

func SearchProduct() gin.HandlerFunc {}

func SearchProductByQuery() gin.HandlerFunc {}