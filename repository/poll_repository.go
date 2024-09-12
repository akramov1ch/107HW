package repository

import (
    "errors"
    "107HW/models"
    "sync"
)

var (
    polls = make(map[string]*models.Poll) 
    mu    = &sync.Mutex{}
)

func CreatePoll(poll *models.Poll) error {
    mu.Lock()
    defer mu.Unlock()

    if _, exists := polls[poll.ID]; exists {
        return errors.New("Poll already exists")
    }

    poll.Votes = make(map[string]int)
    polls[poll.ID] = poll
    return nil
}

func VotePoll(pollID string, userIP string, vote *models.Vote) error {
    mu.Lock()
    defer mu.Unlock()

    poll, exists := polls[pollID]
    if !exists {
        return errors.New("Poll not found")
    }

    if _, voted := poll.Votes[userIP]; voted {
        return errors.New("User has already voted")
    }

    poll.Votes[userIP] = 1
    return nil
}

func GetPollResults(pollID string) (*models.Poll, error) {
    mu.Lock()
    defer mu.Unlock()

    poll, exists := polls[pollID]
    if !exists {
        return nil, errors.New("Poll not found")
    }

    return poll, nil
}
