package main

import (
	"fmt"
	"platform-game/model"
)

func main() {
	users := []model.User{{ID: 1, Username: "John", Active: true}, {ID: 2, Username: "Sarah", Active: true}, {ID: 1, Username: "", Active: true}}
	fmt.Println("Lista de usuários válidos:")
	for _, user := range users {
	if !user.IsValid() {
		continue}
		fmt.Println(user)
	}}