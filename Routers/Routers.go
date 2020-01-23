package Routers

import (
	"net/http"
	"../Controllers"
	"../Models"
	"fmt"
	"github.com/gin-gonic/gin"
	jwt "github.com/dgrijalva/jwt-go"
)
func loginHandler(c *gin.Context) {
	var user Models.Credential
	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "can't bind struct",
		})
	}
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username != "myname" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "wrong username",
		})
		c.Abort()
	} else if password != "myname123" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "wrong password",
		})
		c.Abort()
	}else{
		sign := jwt.New(jwt.GetSigningMethod("HS256"))
		token, err := sign.SignedString([]byte("secret"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			c.Abort()
		}
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})	
	}
	
}

func auth(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		
		return []byte("secret"), nil
	})
	
	if token != nil && err == nil {
		fmt.Println("token verified")
		} else {
			result := gin.H{
				"message": "not authorized",
				"error":   err.Error(),
			}
			c.JSON(http.StatusUnauthorized, result)
			c.Abort()
		}
	}
	
	func SetupRouter() *gin.Engine {
		r := gin.Default()
		
		v1 := r.Group("/v1")
		{
			v1.POST("login", loginHandler)
			v1.GET("book", auth, Controllers.ListBook)
			v1.POST("book", Controllers.AddNewBook)
			v1.GET("book/:id", Controllers.GetOneBook)
			v1.PUT("book/:id", Controllers.PutOneBook)
			v1.DELETE("book/:id", Controllers.DeleteBook)
		}
		
		v2 := r.Group("/v2")
		{
			v2.GET("author", Controllers.ListAuthor)
		}
		
		return r
	}
	