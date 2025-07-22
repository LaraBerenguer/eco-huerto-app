package main

import (
	"database/sql"
	"ecohortapp/repository"
	"log"
	"net/http"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
	_ "github.com/glebarez/go-sqlite"
)

type Config struct {
	App                                       fyne.App              //base para construir la GUI (visual)
	InfoLog                                   *log.Logger           //log de acciones del usuario
	ErrorLog                                  *log.Logger           //log de errores
	DB                                        repository.Repository //puntero de conexión a la db
	MainWindow                                fyne.Window           //ventana principal con fyne
	ClimaDadesContainer                       *fyne.Container       //almacenar contenedor de los datos aemet
	HTTPClient                                http.Client           //definir carga de la librería http
	ForecastGraphContainer                    *fyne.Container       //contenedor para guardar el gráfico
	RegistresTable                            *widget.Table         //tabla para guardar datos de db
	Registres                                 [][]interface{}       //guardar slice de slice obtenido de la db
	AddRegistresDataRegistresEntrada          *widget.Entry         //guardar campo data del formulario
	AddRegistresPrecipitacionRegistresEntrada *widget.Entry         //guardar campo precipitación del formulario
	AddRegistresTempMaximaRegistresEntrada    *widget.Entry         //guardar campo temp max del formulario
	AddRegistresTempMinimaRegistresEntrada    *widget.Entry         //guardar campo temp min del formulario
	AddRegistresHumedadRegistresEntrada       *widget.Entry         //guardar campo humedad del formulario
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
	sqlDB, err := myApp.ConnectSQL()
	if err != nil {
		log.Panic(err)
	}

	//Crear el repositorio de la base de datos
	myApp.setupDB(sqlDB)

	//Definir tamaño y otras características de la ventana
	myApp.MainWindow = laMevaApp.NewWindow("Eco Hort App")
	myApp.MainWindow.Resize(fyne.NewSize(800, 700)) //tamaño de la ventana
	myApp.MainWindow.SetFixedSize(true)             //tamaño fijo
	myApp.MainWindow.SetMaster()                    //pantalla principal, si cierra esto cierra todo

	myApp.makeUI() //llamada a la función para generar la UI

	//Ejecutar la app
	myApp.MainWindow.ShowAndRun()
}

func (myApp *Config) ConnectSQL() (*sql.DB, error) {
	path := ""

	//detecta si tenemos variable de entorno de db
	if os.Getenv("DBpath") != "" {
		path = os.Getenv("DBpath")
	} else {
		path = myApp.App.Storage().RootURI().Path() + "/sql.db"
		//log
		myApp.InfoLog.Println("La base de datos está en: ", path)
	}

	con, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	return con, nil
}

//ejecutar e instalar estructuras

func (myApp *Config) setupDB(sqlDB *sql.DB) {
	myApp.DB = repository.NewSQLiteRepository(sqlDB)

	err := myApp.DB.Migrate()
	if err != nil {
		myApp.ErrorLog.Println(err)
		log.Panic()
	}
}
