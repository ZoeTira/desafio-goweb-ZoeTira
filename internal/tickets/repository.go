package tickets

import (
	//	"fmt"
	"fmt"

	"github.com/ZoeTira/desafio-goweb-ZoeTira/domain"
	"github.com/ZoeTira/desafio-goweb-ZoeTira/pkg/store"
)
//var ticketList []domain.Ticket

type Repository interface{
	GetAll()([]domain.Ticket, error)
	GetTicketByDestination(destination string)([]domain.Ticket, error)
	GetCountTicketsByDestination(destination string)(int, error)
	GetAVGTicketByDestination(destination string)(float64, error)
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
	ticketList, err := r.db.Read()
	if err != nil {
		return []domain.Ticket{}, err
	}
//Verificar que no sea nula la list

	//Retorno todos los tickets
	return ticketList, nil

}

func (r *repository) GetTicketByDestination(destination string)([]domain.Ticket, error){
	ticketList, err := r.db.Read()
	if err != nil {
		return []domain.Ticket{}, err
	}
	var ticketsByCountry []domain.Ticket
	for _, t := range ticketList{
		if(t.Country == destination){
			ticketsByCountry = append(ticketsByCountry,t)
		}
	}
	
	return ticketsByCountry, nil
}

func (r *repository) GetCountTicketsByDestination(destination string)(int, error){
	ticketsList, err := r.GetTicketByDestination(destination)
	if(err != nil){
		return 0, err
	}
	if(len(ticketsList) == 0){
		return 0, fmt.Errorf("empty ticket list")
	}

	return len(ticketsList),nil
	
}

func (r *repository) GetAVGTicketByDestination(destination string)(float64, error){
	ticketList, err := r.db.Read()
	if err != nil {
		return 0, err
	}

	var ticketsByCountry []domain.Ticket
	for _, t := range ticketList{
		if(t.Country == destination){
			ticketsByCountry = append(ticketsByCountry,t)
		}
	}
	
	avg := float64((len(ticketsByCountry) * 100.0) / (len(ticketList)))
	return avg, nil
}