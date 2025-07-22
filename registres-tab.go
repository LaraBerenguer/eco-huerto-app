package main

import (
	"ecohortapp/repository"
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func (myApp *Config) registresTab() *fyne.Container {
	myApp.Registres = myApp.getRegistresSlice()
	myApp.RegistresTable = myApp.getRegistresTable()
	contenidorRegistres := container.NewBorder(
		nil,
		nil,
		nil,
		nil,
		container.NewAdaptiveGrid(1, myApp.RegistresTable),
	)
	return contenidorRegistres
}

func (myApp *Config) getRegistresTable() *widget.Table {

	//tabla
	t := widget.NewTable(
		func() (int, int) {
			return len(myApp.Registres), len(myApp.Registres[0])
		},
		func() fyne.CanvasObject {
			ctn := container.NewVBox(widget.NewLabel(""))
			return ctn
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			if i.Col == (len(myApp.Registres[0])-1) && i.Row != 0 { //si la celda es la última col y no primera fila
				w := widget.NewButtonWithIcon("Borrar", theme.DeleteIcon(), func() {
					//ventana de confirmar
					dialog.ShowConfirm("¿Seguro que quieres borrar este elemento?", "", func(deleted bool) {
						if deleted {
							id, _ := strconv.Atoi(myApp.Registres[i.Row][0].(string))
							err := myApp.DB.BorrarRegistro(int64(id))
							if err != nil {
								myApp.ErrorLog.Println(err)
							}
						}
						myApp.actualitzarRegistresTab()

					}, myApp.MainWindow)
				})

				w.Importance = widget.HighImportance
				o.(*fyne.Container).Objects = []fyne.CanvasObject{
					w,
				}
			} else {
				o.(*fyne.Container).Objects = []fyne.CanvasObject{
					widget.NewLabel(myApp.Registres[i.Row][i.Col].(string)),
				}
			}
		})

	ampleColumne := []float32{50, 115, 115, 115, 115, 115, 120}
	for i := 0; i < len(ampleColumne); i++ {
		t.SetColumnWidth(i, ampleColumne[i])
	}
	return t
}

// slice de slices
func (myApp *Config) getRegistresSlice() [][]interface{} {
	var slice [][]interface{}
	registres, err := myApp.registresActuals()
	//control de error
	if err != nil {
		myApp.ErrorLog.Println(err)
	}

	slice = append(slice, []interface{}{"ID", "Data", "Precipitacio", "Temp. Máxima", "Temp. Mínima", "Humedad"})

	for _, v := range registres {
		var registre []interface{}

		registre = append(registre, strconv.FormatInt(v.ID, 10))
		registre = append(registre, v.Data.Format("2006-01-02"))
		registre = append(registre, fmt.Sprintf("%d%%", v.Precipitacio))
		registre = append(registre, fmt.Sprintf("%d", v.TempMaxima))
		registre = append(registre, fmt.Sprintf("%d", v.TempMinima))
		registre = append(registre, fmt.Sprintf("%d%%", v.Humitat))
		registre = append(registre, widget.NewButton("Borrar", func() {}))
		slice = append(slice, registre)
	}
	//gestion de datos
	return slice
}

// obtener slice de db
func (myApp *Config) registresActuals() ([]repository.Registros, error) {
	registres, err := myApp.DB.LeerRegistros()
	if err != nil {
		myApp.ErrorLog.Println(err)
		return nil, err
	}

	return registres, nil
}
