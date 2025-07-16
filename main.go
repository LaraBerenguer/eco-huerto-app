package main

import (
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	//"fyne.io/fyne/v2/widget"
)

type Config struct {
	App        fyne.App    //base para construir la GUI (visual)
	InfoLog    *log.Logger //log de acciones del usuario
	ErrorLog   *log.Logger //log de errores
	MainWindow fyne.Window //ventana principal con fyne
}

var myApp Config

func main() {
	//Desarrollo de la base de datos
	laMevaApp := app.NewWithID("es.huertourbano.desktop")
	myApp.App = laMevaApp

	//Definir logs
	myApp.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	myApp.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Lshortfile)

	//Conexión con la base de datos
	//Crear el repositorio de la base de datos

	//Definir tamaño y otras características de la ventana
	myApp.MainWindow = laMevaApp.NewWindow("Eco Hort App")
	myApp.MainWindow.Resize(fyne.NewSize(800, 500)) //tamaño de la ventana
	myApp.MainWindow.SetFixedSize(true)             //tamaño fijo
	myApp.MainWindow.SetMaster()                    //pantalla principal, si cierra esto cierra todo

	myApp.makeUI() //llamada a la función para generar la UI
	//Ejecutar la app
	myApp.MainWindow.ShowAndRun()
}
