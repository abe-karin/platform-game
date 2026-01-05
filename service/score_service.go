package service

import "platform-game/model"
func CalculateTotalScore(scores []model.Score, userId int) int {
	total := 0
	for _, score := range scores {
		if score.UserID == userId {
			total += score.Points
		}
	}
	return total
}