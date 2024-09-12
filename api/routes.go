package api

import (
    "107HW/controllers"
    "github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
    pollRoutes := r.Group("/polls")
    {
        pollRoutes.POST("/", controllers.CreatePoll)       
        pollRoutes.GET("/:id", controllers.GetPollResults) 
        pollRoutes.POST("/:id/vote", controllers.VotePoll) 
    }
}
