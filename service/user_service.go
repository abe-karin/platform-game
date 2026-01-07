package service

import "platform-game/model"

func FilterValidUsers(users []model.User) []model.User {
	validUsers := make([]model.User, 0)
	for _, user := range users {
		if err := user.Validate(); err != nil {
			continue
		}
		if !user.IsValid() {
			continue
		}
		validUsers = append(validUsers, user)
	}
	return validUsers
}
