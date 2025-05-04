package service

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"minesweeper/constant"
)

type Game struct {
	app        fyne.App
	Window     fyne.Window
	containers map[string]*fyne.Container
	Mine       *Mine
}

func NewGame() *Game {
	game := new(Game)
	game.app = app.New()
	game.Window = game.app.NewWindow("扫雷")
	return game
}

func (g *Game) CreateStartInterface() {
	gameLevel := widget.NewLabel("游戏难度")
	selectLevel := widget.NewSelect([]string{constant.Simple, constant.Common},
		func(s string) {
			g.Mine = mines[s]
		})
	selectLevel.SetSelected(constant.Common)
	level := container.NewHBox(gameLevel, selectLevel)

	startButton := widget.NewButton("开始游戏", func() {
		g.CreateGameInterface()
	})

	g.containers[constant.StartInterface] = container.NewCenter(container.NewVBox(level, startButton))

	g.Window.SetContent(g.containers[constant.StartInterface])
}

func (g *Game) CreateGameInterface() {
	content := container.NewGridWithColumns(g.Mine.col)
	g.Mine.Run()
	for _, c := range g.Mine.Cells {
		content.Add(c.button)
	}
	g.Window.SetContent(content)
	g.Window.Resize(fyne.NewSize(float32(40*g.Mine.row), float32(40*g.Mine.col)))
	g.containers[constant.GameInterface] = content
}

func (g *Game) Run() {
	g.containers = make(map[string]*fyne.Container)
	g.CreateStartInterface()
	g.Window.Resize(fyne.NewSize(400, 400))
	g.Window.ShowAndRun()
}
