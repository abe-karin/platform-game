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
	fmt.Println("Valid Users:")
	scores := []model.Score{
		{UserID: 1, Points: 10},
		{UserID: 2, Points: 20},
		{UserID: 1, Points: 15},
	}

	for _, user := range validUsers {
		total := service.CalculateTotalScore(scores, user.ID)
		fmt.Printf("Total score for user %s: %d\n", user.Username, total)
	}

	service.ProcessWithDefer()
	ranking := service.BuildRanking(scores)
	fmt.Println("Ranking:")
	for i, entry := range ranking {
		fmt.Printf("%d lugar - User %d: %d pontos\n", i+1, entry.UserID, entry.Points)
	}
}
