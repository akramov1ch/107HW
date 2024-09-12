package main

import (
    "107HW/api"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    api.SetupRoutes(r)

    r.Run(":8080")
}
