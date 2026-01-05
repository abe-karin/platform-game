package main

import (
	"fmt"
	"platform-game/model"
	"platform-game/service"
)

func main() {
    users := []model.User{
        {ID: 1, Username: "John", Active: true}, 
        {ID: 2, Username: "Sarah", Active: true}, 
        {ID: 3, Username: "Sebastian", Active: true},
		{ID: 4, Username: "", Active: true},
    }
validUsers := service.FilterValidUsers(users) 
    
   

	scores := []model.Score{
		{UserID: 1, Points: 10},
		{UserID: 2, Points: 20},
		{UserID: 3, Points: 15},
	}

	for _, user := range validUsers {
		total := service.CalculateTotalScore(scores, user.ID)
		fmt.Printf("Total score for user %s: %d\n", user.Username, total)
	}
}


    