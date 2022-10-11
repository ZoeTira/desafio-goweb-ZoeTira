package tickets

import (
//	"fmt"
	"github.com/ZoeTira/desafio-goweb-ZoeTira/pkg/store"
	"github.com/ZoeTira/desafio-goweb-ZoeTira/domain"


)
var ticketList []domain.Ticket

type Repository interface{
	GetAll()([]domain.Ticket, error)
	GetTicketByDestination(destination string)([]domain.Ticket, error)
}

type repository struct{
	db store.Store
}

func NewRepository(db store.Store) Repository{
	return &repository{
		/*aca iria bd*/
		db: db,
	}
}

func (r *repository) GetAll()([]domain.Ticket, error){
	//Leo y le paso donde guardar
	err := r.db.Read(&ticketList)
	if err != nil {
		return []domain.Ticket{}, err
	}
//Verificar que no sea nula la list

	//Retorno todos los tickets
	return ticketList, nil

}

func (r *repository) GetTicketByDestination(destination string)([]domain.Ticket, error){
	err := r.db.Read(&ticketList)
	if err != nil {
		return []domain.Ticket{}, err
	}
	
	return ticketList, nil
}