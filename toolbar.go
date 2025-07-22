package main

import (
	"ecohortapp/repository"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func (myApp *Config) getToolbar(window fyne.Window) *widget.Toolbar {
	toolbar := widget.NewToolbar(
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			myApp.addRegistresDialog()
		}),
		widget.NewToolbarAction(theme.ViewRefreshIcon(), func() {
			myApp.actualitzarClimadadesContent()
		}),
		widget.NewToolbarAction(theme.SettingsIcon(), func() {}),
	)
	return toolbar
}

func (myApp *Config) addRegistresDialog() dialog.Dialog {
	dataRegistreEntrada := widget.NewEntry()
	precipitacionRegistreEntrada := widget.NewEntry()
	tempMaximaRegistreEntrada := widget.NewEntry()
	tempMinimaRegistreEntrada := widget.NewEntry()
	humedadRegistreEntrada := widget.NewEntry()

	myApp.AddRegistresDataRegistresEntrada = dataRegistreEntrada
	myApp.AddRegistresPrecipitacionRegistresEntrada = precipitacionRegistreEntrada
	myApp.AddRegistresTempMaximaRegistresEntrada = tempMaximaRegistreEntrada
	myApp.AddRegistresTempMinimaRegistresEntrada = tempMinimaRegistreEntrada
	myApp.AddRegistresHumedadRegistresEntrada = humedadRegistreEntrada

	validarData := func(text string) error {
		if _, err := time.Parse("2006-01-02", text); err != nil {
			return err
		}
		return nil
	}

	dataRegistreEntrada.Validator = validarData
	dataRegistreEntrada.PlaceHolder = "YYY-MM-DD"

	//validacion enteros
	validarInt := func(text string) error {
		_, err := strconv.Atoi(text)
		if err != nil {
			return err
		}
		return nil
	}

	precipitacionRegistreEntrada.Validator = validarInt
	tempMaximaRegistreEntrada.Validator = validarInt
	tempMinimaRegistreEntrada.Validator = validarInt
	humedadRegistreEntrada.Validator = validarInt

	addForm := dialog.NewForm(
		"Nuevo Registro",
		"Guardar",
		"Cancelar",
		[]*widget.FormItem{
			{Text: "Fecha de registro", Widget: dataRegistreEntrada},
			{Text: "Probabilidad de Precipitación", Widget: precipitacionRegistreEntrada},
			{Text: "Temperatura Máxima", Widget: tempMaximaRegistreEntrada},
			{Text: "Temperatura Mínima", Widget: tempMinimaRegistreEntrada},
			{Text: "Humedad", Widget: humedadRegistreEntrada},
		},
		func(valid bool) {
			if valid {
				dataRegistre, _ := time.Parse("2006-01-02", dataRegistreEntrada.Text)
				precipitacion, _ := strconv.Atoi(precipitacionRegistreEntrada.Text)
				tempMaxima, _ := strconv.Atoi(tempMaximaRegistreEntrada.Text)
				tempMinima, _ := strconv.Atoi(tempMinimaRegistreEntrada.Text)
				humedad, _ := strconv.Atoi(humedadRegistreEntrada.Text)

				//insertar en bd
				_, err := myApp.DB.InsertarRegistro(repository.Registros{
					Data:         dataRegistre,
					Precipitacio: precipitacion,
					TempMaxima:   tempMaxima,
					TempMinima:   tempMinima,
					Humitat:      humedad,
				})
				if err != nil {
					myApp.ErrorLog.Println(err)
				}
				myApp.actualitzarRegistresTab()
			}
		},
		myApp.MainWindow)

	addForm.Resize(fyne.Size{Width: 400})
	addForm.Show()
	return addForm
}
