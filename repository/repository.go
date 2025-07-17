package repository

import (
	"errors"
	"time"
)

var (
	errorActualizantDades = errors.New("error actualizando datos")
	errorBorrantDades     = errors.New("error borrando datos")
)

type Repositoty interface {
	Migrate() error
	InsertarRegistro(nuevoRegistro Registros) (*Registros, error)
	LeerRegistro(id int64) (*Registros, error)
	LeerRegistros() ([]Registros, error)
	ActualizarRegistro(id int64, actualizar *Registros) error
	BorrarRegistro(id int64) error
}

type Registros struct {
	ID           int64     `json:"id"`
	Data         time.Time `json:"data_registre"`
	Precipitacio int       `json:"precipitacio"`
	TempMaxima   int       `json:"temp_maxima"`
	TempMinima   int       `json:"temp_minima"`
	Humitat      int       `json:"humedad"`
}
