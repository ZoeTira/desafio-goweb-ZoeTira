package tickets

import (
	"fmt"
	"github.com/ZoeTira/desafio-goweb-ZoeTira/domain"

)

type Repository interface{
	GetAll()([]domain.Ticket, error)
	GetTicketByDestination(destination string)([]domain.Ticket)
}

type repository struct{
	//Puedo agregarle la bd
}

func NewRepository(/*aca iria bd*/) Repository{
	return &repository{
		/*aca iria bd*/
	}
}

func (r *repository)GetAll()([]domain.Ticket, error){
	var ps []domain.Ticket
	//Leo t le paso donde guardar
	err := r.db.Read(&ps)
	if err != nil {
		return []Product{}, err
	}
	return ps, nil

}