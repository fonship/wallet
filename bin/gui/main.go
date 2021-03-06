// Package main provides various examples of Fyne API capabilities
package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"
	"github.com/lengzhao/wallet/bin/gui/event"
	"github.com/lengzhao/wallet/bin/gui/res"
	"github.com/lengzhao/wallet/bin/gui/screens"
)

func main() {
	res.LoadLanguage()
	a := app.NewWithID("govm.net")
	a.SetIcon(res.GetResource("resource/img/govm.png"))
	a.Settings().SetTheme(theme.LightTheme())

	event.Send(event.ELogin)
	// screens.Login(a)
	screens.Master(a)
	a.Run()
}
