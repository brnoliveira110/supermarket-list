package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    // Inicializa o Gin
    r := gin.Default()

    // Rota de teste
    r.GET("/api/health", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Backend Go está funcionando!",
        })
    })

    // Roda o servidor na porta 8080
    r.Run(":8080")
}