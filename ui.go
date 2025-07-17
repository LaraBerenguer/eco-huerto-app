package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

func (myApp *Config) makeUI() {
	//Conectar con AEMET
	precipitacio, tempMax, tempMin, humitat := myApp.getClimaText()

	//Formatear e insertar datos en un contenedor
	contenidorClima := container.NewGridWithColumns(4,
		precipitacio,
		tempMax,
		tempMin,
		humitat)

	//Meter contenedor pequeño dentro de contenedor ventana
	myApp.ClimaDadesContainer = contenidorClima

	//toolbar
	toolbar := myApp.getToolbar(myApp.MainWindow)

	//grafico primera pestaña
	contenidorPrimeraTab := myApp.forecastTab()

	//pestañas
	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Previsión", theme.HomeIcon(), canvas.NewText("Texto de ejemplo", nil)),
		container.NewTabItemWithIcon("Parte Meteorológico", theme.ComputerIcon(), canvas.NewText("Otro texto de ejemplo", nil)),
	)

	tabs.SetTabLocation(container.TabLocationTop)

	//situar el contenedor en la ventana
	verticalContainer := container.NewVBox(contenidorClima, toolbar, tabs, contenidorPrimeraTab)
	myApp.MainWindow.SetContent(verticalContainer)

	//crear go rutine
	go func() {
		for range time.Tick(time.Second * 30) {
			myApp.actualitzarClimadadesContent()
		}
	}()
}

func (myApp *Config) actualitzarClimadadesContent() {
	myApp.InfoLog.Println("refrescando valores meteo") //log
	precipitacio, tempMax, tempMin, humitat := myApp.getClimaText()
	myApp.ClimaDadesContainer.Objects = []fyne.CanvasObject{precipitacio, tempMax, tempMin, humitat}
	myApp.ClimaDadesContainer.Refresh()

	image := myApp.getImg()
	myApp.ForecastGraphContainer.Objects = []fyne.CanvasObject{image}
	myApp.ForecastGraphContainer.Refresh()
}
