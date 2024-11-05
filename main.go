package main

import (
	"net/http"
	"regexp"
	"strconv"

	"go-api/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	magic := r.Group("/magic")
	{

		magic.GET("/sum", func(c *gin.Context) {
			n, _ := strconv.Atoi(c.DefaultQuery("n", "0"))
			c.JSON(http.StatusOK, gin.H{
				"result": utils.MagicSum(n),
			})
		})

		magic.POST("/pow", func(c *gin.Context) {
			var data struct {
				n int `json:"n" binding:"required"`
			}
			if err := c.ShouldBindJSON(&data); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"result": utils.MagicPow(data.n),
			})
		})

		magic.GET("/odd", func(c *gin.Context) {
			n, _ := strconv.Atoi(c.DefaultQuery("n", "0"))
			c.JSON(http.StatusOK, gin.H{
				"result": utils.Magicodd(n),
			})
		})

		magic.POST("/grade", func(c *gin.Context) {
			var data struct {
				n int `json:"n" binding:"required"`
			}
			if err := c.ShouldBindJSON(&data); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"result": utils.MagicGrade(data.n),
			})
		})

		magic.GET("/name", func(c *gin.Context) {
			n, _ := strconv.Atoi(c.DefaultQuery("n", "0"))
			c.JSON(http.StatusOK, gin.H{
				"result": utils.MagicName(n),
			})
		})

		magic.POST("/tria", func(c *gin.Context) {
			var data struct {
				n int `json:"n" binding:"required"`
			}
			if err := c.ShouldBindJSON(&data); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"result": utils.MagicTria(data.n),
			})
		})
	}

	account := r.Group("/account")
	{
		account.GET("/create", createAccountHandler)
		account.GET("/read", readAccountHandler)
		account.GET("/update", updateAccountHandler)
		account.GET("/delete", deleteAccountHandler)
		account.GET("/list", listAccountHandler)

		// Replace "yourusername" with your actual username
		account.GET("/username", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"data": "Hi, my name is joyce"})
		})
	}

	r.POST("/auth/login", loginAuthHandler)

	r.Run(":8081")
}

func createAccountHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}

func readAccountHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": gin.H{}})
}

func updateAccountHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}

func deleteAccountHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}

func listAccountHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": []interface{}{}})
}

// Bonus: Authentication handler
func loginAuthHandler(c *gin.Context) {
	var loginData struct {
		Username string `json:"username" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	// Bind JSON data from the request to loginData struct
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if username is alphanumeric and password is numeric
	isUsernameValid := regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString(loginData.Username)
	isPasswordValid := regexp.MustCompile(`^[0-9]+$`).MatchString(loginData.Password)

	if isUsernameValid && isPasswordValid {
		c.JSON(http.StatusOK, gin.H{"message": "Login success"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Login failed"})
	}
}
