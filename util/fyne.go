package util

import "fyne.io/fyne/v2"

func GetTopWindow() fyne.Window {
	return fyne.CurrentApp().Driver().AllWindows()[0]
}
