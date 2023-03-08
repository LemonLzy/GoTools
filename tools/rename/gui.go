package tools

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"path/filepath"
)

func Gui() {
	a := app.NewWithID("com.xxx.yyy")

	absPath, _ := filepath.Abs("BatchRename/png/logo.jpg")
	fmt.Println(absPath)
	logo, _ := fyne.LoadResourceFromPath("BatchRename/png/logo.jpg")
	a.SetIcon(logo)

	w := a.NewWindow("Hello")
	w.Resize(fyne.NewSize(249, 100))

	hello := widget.NewLabel("Hello Fyne!")
	a.Preferences().SetString("a", "preference value")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText(a.Preferences().String("a"))
		}),
	))

	w.ShowAndRun()
}
