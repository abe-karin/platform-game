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
        {ID: 1, Username: "", Active: true},
    }
    
    validUsers := service.FilterValidUsers(users) 
    
    for _, user := range validUsers {
        fmt.Println(user)
    }
}