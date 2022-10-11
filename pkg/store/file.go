package store

import (
	"encoding/csv"
	"os"
	"fmt"

	"strconv"
	"github.com/ZoeTira/desafio-goweb-ZoeTira/domain"

)

// Implemento interface store, metodos leer y escribir, les paso una interfaz
type Store interface {
	Read() ([]domain.Ticket, error)
	//Write(date interface{}) error
}

// Es como un ""Alias"" de string, me va a servir desp para comparar y que no me
// rompa todo. Es para que desp depeneidno del tipo de archivo que tenga
// el new me lo construya
type Type string

type fileStore struct {
	FilePath string
}

const (
	FileType Type = "file"
)

func New(store Type, fileName string) Store {
	switch store {
	case FileType:
		return &fileStore{fileName}
	}
	return nil
}

func (f *fileStore) Read() ([]domain.Ticket, error){

//genero variable donde voy a guardar los tickets
	var ticketList []domain.Ticket
//Abro el archivo
	file, err := os.Open(f.FilePath)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}
//Leo y guardo en csvR la data, que lea todo
	csvR := csv.NewReader(file)
	data, err := csvR.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}
//Esta todo guardado en data, voy leyendo fila por fila extrayendo los datos y armando mi estructura de ticket
	for _, row := range data{
		price, err := strconv.ParseFloat(row[5], 64)
		if err != nil {
			return nil, err
		}
		ticketList = append(ticketList, domain.Ticket{
			Id:      row[0],
			Name:    row[1],
			Email:   row[2],
			Country: row[3],
			Time:    row[4],
			Price:   price,
		})
		
	}

	return ticketList, nil
}


/* Metodo para escribir. No hecho
func (f *fileStore) Write(data interface{}) error{
	file, err := json.MarshalIndent(data, "", " ")
    if err!= nil {
        return err
    }

	return os.WriteFile(f.FilePath, file, 0644)
}*/
