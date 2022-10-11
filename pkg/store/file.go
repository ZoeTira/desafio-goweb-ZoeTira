package store


import(
	"encoding/json"
	"os"
)

//Implemento interface store, metodos leer y escribir, les paso una interfaz
type Store interface {
    Read(data interface{}) error
	//Write(date interface{}) error
}

//Es como un ""Alias"" de string, me va a servir desp para comparar y que no me 
//rompa todo. Es para que desp depeneidno del tipo de archivo que tenga
//el new me lo construya
type Type string

type fileStore struct{
	FilePath string
}

const(
	FileType Type = "file"
)

func New(store Type, fileName string) Store{
	switch store {
    case FileType:
        return &fileStore{fileName}
	}
	return nil
}

func (f *fileStore) Read(data interface{}) error{
	file, err := os.ReadFile(f.FilePath)
    if err!= nil {
        return err
    }
	return json.Unmarshal(file, &data)
}
/*
func (f *fileStore) Write(data interface{}) error{
	file, err := json.MarshalIndent(data, "", " ")
    if err!= nil {
        return err
    }

	return os.WriteFile(f.FilePath, file, 0644)
}*/