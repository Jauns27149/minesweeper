package main

import (
	"log"
	"minesweeper/service"
)

func main() {
	game := service.NewGame()
	log.Println("minesweeper start...")
	game.Run()
}
