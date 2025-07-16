package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func (myApp *Config) getClimaText() (*canvas.Text, *canvas.Text, *canvas.Text, *canvas.Text) {
	var parte Diaria
	var precipitacio, tempMax, tempMin, humitat *canvas.Text
	prediccion, err := parte.getPredictions()

	if err != nil {
		///connection error
		gris := color.NRGBA{R: 155, G: 155, B: 155, A: 255}
		precipitacio = canvas.NewText("Precipitación: Undefined", gris)
		tempMax = canvas.NewText("Temp Max: Undefined", gris)
		tempMin = canvas.NewText("Temp Min: Undefined", gris)
		humitat = canvas.NewText("Humedad: Undefined", gris)
	} else {
		colorTexte := color.NRGBA{R: 0, G: 180, B: 0, A: 255} //verde
		if prediccion.ProbPrecipitacio < 50 {
			colorTexte = color.NRGBA{R: 180, G: 0, B: 0, A: 255} //rojo
		}

		//Preparar strings
		precipitacionTtx := fmt.Sprintln("Precipitación: %d%%", prediccion.ProbPrecipitacio)
		tempMaxTxt := fmt.Sprintln("Temp. Máxima: %d", prediccion.TemperaturaMax)
		tempMinTxt := fmt.Sprintln("Temp. Mínima: %d", prediccion.TemperaturaMin)
		humedadTxt := fmt.Sprintln("Humedad Relativa: %d%%", prediccion.HumitatRelativa)

		//crear elementos de texto
		precipitacio = canvas.NewText(precipitacionTtx, colorTexte)
		tempMax = canvas.NewText(tempMaxTxt, nil)
		tempMin = canvas.NewText(tempMinTxt, nil)
		humitat = canvas.NewText(humedadTxt, nil)
	}

	precipitacio.Alignment = fyne.TextAlignLeading
	tempMax.Alignment = fyne.TextAlignCenter
	tempMin.Alignment = fyne.TextAlignCenter
	humitat.Alignment = fyne.TextAlignTrailing

	return precipitacio, tempMax, tempMin, humitat
}
