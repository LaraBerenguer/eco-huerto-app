package main

import (
	"bytes"
	"errors"
	"image"
	"image/png"
	"io"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func (app *Config) forecastTab() *fyne.Container {
	image := myApp.getImg()
	containerImg := container.NewVBox(image)
	myApp.ForecastGraphContainer = containerImg
	return containerImg
}

func (app *Config) getImg() *canvas.Image {
	url := "https://my.meteoblue.com/visimage/meteogram_web_hd?look=KILOMETER_PER_HOUR%2CCELSIUS%2CMILLIMETER&apikey=5838a18e295d&temperature=C&windspeed=kmh&precipitationamount=mm&winddirection=3char&city=Abrera&iso2=es&lat=41.5168&lon=1.901&asl=111&tz=Europe%2FMadrid&lang=es&sig=b353aab637f77ab97ae54cbd760554f2"
	var img *canvas.Image
	err := myApp.downloadFile(url, "forecast.png")

	if err != nil {
		img = canvas.NewImageFromResource(resourceNodisponiblePng)
	} else {
		img = canvas.NewImageFromFile("forecast.png")
	}

	img.SetMinSize(fyne.Size{
		Width:  770,
		Height: 410,
	})

	return img
}

func (app *Config) downloadFile(url, fileName string) error {
	res, err := myApp.HTTPClient.Get(url)
	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		return errors.New("ha habido un error")
	}

	bytesSlice, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	//Decodificar imagen a slice de bytes
	img, _, err := image.Decode(bytes.NewReader(bytesSlice))
	if err != nil {
		return err
	}

	//crear archivo y guardar imagen en archivo
	file, err := os.Create("./" + fileName)
	if err != nil {
		return err
	}

	err = png.Encode(file, img) //?
	if err != nil {
		return err
	}

	return nil
}
