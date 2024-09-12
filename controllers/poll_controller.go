package controllers

import (
    "net/http"
    "107HW/models"
    "107HW/repository"
    "github.com/gin-gonic/gin"
)

func CreatePoll(c *gin.Context) {
    var poll models.Poll
    if err := c.ShouldBindJSON(&poll); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    if err := repository.CreatePoll(&poll); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create poll"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Poll created successfully"})
}

func VotePoll(c *gin.Context) {
    pollID := c.Param("id")
    var vote models.Vote
    if err := c.ShouldBindJSON(&vote); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    ip := c.ClientIP() 
    if err := repository.VotePoll(pollID, ip, &vote); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Vote registered successfully"})
}

func GetPollResults(c *gin.Context) {
    pollID := c.Param("id")
    
    pollResults, err := repository.GetPollResults(pollID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get poll results"})
        return
    }

    c.JSON(http.StatusOK, pollResults)
}
