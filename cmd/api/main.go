package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/marriosdev/export-api/internal/api"
)

func main() {
	r := gin.Default()

	api.RegisterRouter(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("não foi possível iniciar servidor: %v", err)
	}
}
