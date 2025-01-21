package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    // Inicializa o router utilizando as configs default do Gin
    r := gin.Default()

    // Definindo uma rota com metodo GET para o endpoint /ping e passa uma function como parametro (nosso handler)
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })

    // Roda a api
    r.Run()
}