package repository

import (
	"database/sql"
	"time"
)

type SQLiteRepository struct {
	Conn *sql.DB
}

func NewSQLiteRepository(db *sql.DB) *SQLiteRepository {
	return &SQLiteRepository{
		Conn: db,
	}
}

func (repo *SQLiteRepository) Migrate() error {
	sentencia := `create table if not exists registres( 
	id integer primary key autoincrement,
	data_registre integer not null,
	precipitacio integer not null,
	temp_maxima integer not null,
	temp_minima integer not null,
	humitat integer not null)`

	_, err := repo.Conn.Exec(sentencia)
	return err
}

func (repo *SQLiteRepository) InsertarRegistro(nuevoRegistro Registros) (*Registros, error) {
	query := "insert into registres (data_registre,precipitacio,temp_maxima,temp_minima,humitat) values (?,?,?,?,?)"
	res, err := repo.Conn.Exec(query, nuevoRegistro.Data.Unix(), nuevoRegistro.Precipitacio, nuevoRegistro.TempMaxima, nuevoRegistro.TempMinima, nuevoRegistro.Humitat)

	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	nuevoRegistro.ID = id
	return &nuevoRegistro, nil
}

func (repo *SQLiteRepository) LeerRegistros() ([]Registros, error) {
	statement := "select id,data_registre,precipitacio,temp_maxima,temp_minima,humitat from registres order by id desc"
	files, err := repo.Conn.Query(statement)
	if err != nil {
		return nil, err
	}

	defer files.Close()
	var registres []Registros

	for files.Next() { //se puede usar range también pero next se usa para llamadas de bases de datos
		var fila Registros
		var temps int64

		err := files.Scan( //lee cada resultado, va con Next()
			&fila.ID,
			&temps,
			&fila.Precipitacio,
			&fila.TempMaxima,
			&fila.TempMinima,
			&fila.Humitat,
		)

		if err != nil {
			return nil, err
		}

		fila.Data = time.Unix(temps, 0)
		registres = append(registres, fila) //se le mete al slice que hemos creado antes el registro ya poblado a partir de la consulta a la bbdd
	}

	return registres, nil
}

func (repo *SQLiteRepository) LeerRegistro(id int64) (*Registros, error) {
	stmt := "select id,data_registre,precipitacio,temp_maxima,temp_minima,humitat from registres where id = ? order by id desc limit 1"
	filera := repo.Conn.QueryRow(stmt, id) //para cuando solo esperamos una linea de resultados (un resultado)

	var fila Registros
	var temps int64

	err := filera.Scan( //lee el, va con Next()
		&fila.ID,
		&temps,
		&fila.Precipitacio,
		&fila.TempMaxima,
		&fila.TempMinima,
		&fila.Humitat,
	)

	if err != nil {
		return nil, err
	}

	fila.Data = time.Unix(temps, 0)
	return &fila, nil
}
