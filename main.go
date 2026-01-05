package main

import (
	"fmt"
	"platform-game/model"
)

func main() {
	user := model.User{ID: 1, Username: "John", Active: true}
	if !user.IsValid() {
		fmt.Println("Usuário inválido. Operação cancelada.")
		return
	}
	fmt.Println("Usuário válido:", user)
}
