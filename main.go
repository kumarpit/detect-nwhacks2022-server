package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    router := gin.Default()
    router.GET("/", getRoot)

    router.Run("localhost:8000")
}

func getRoot(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, "hello")
}