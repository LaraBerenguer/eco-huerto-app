package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func (myApp *Config) getToolbar(window fyne.Window) *widget.Toolbar {
	toolbar := widget.NewToolbar(
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {}),
		widget.NewToolbarAction(theme.ViewRefreshIcon(), func() {
			myApp.actualitzarClimadadesContent()
		}),
		widget.NewToolbarAction(theme.SettingsIcon(), func() {}),
	)
	return toolbar
}
