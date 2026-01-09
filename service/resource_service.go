package service

import "fmt"
func simulateResource() {
	fmt.Println("Recurso Aberto")
}
func simulateClose() {
	fmt.Println("Recurso Fechado")
}
func ProcessWithDefer() {
	simulateResource()
	defer simulateClose()
	fmt.Println("Processando com Defer")
}