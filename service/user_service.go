package service

import "platform-game/model"

func FilterValidUsers(users []model.User) []model.User { validUsers := []model.User{} 
for _, user := range users { if !user.IsValid() { continue } 
validUsers = append(validUsers, user) } 
return validUsers }
