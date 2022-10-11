package tickets

import (
	"fmt"

	"github.com/ZoeTira/desafio-goweb-ZoeTira/domain"
)

type Service interface{
	GetAll()([]domain.Ticket, error)
	GetTicketByDestination(destination string)([]domain.Ticket, error)
	GetCountTicketsByDestination(destination string)(int, error)
	GetAVGTicketByDestination(destination string)(float64, error)
	
}

type service struct{
	repository Repository
}

func NewService(r Repository) Service{
	return &service{
		repository: r,
	}
}

func (s *service) GetAll()([]domain.Ticket, error){
	ticketList,err := s.repository.GetAll()
	if(err != nil){
		return []domain.Ticket{}, err
	}

	if len(ticketList) == 0{
		return []domain.Ticket{}, fmt.Errorf("empty ticket list")
	}

	return ticketList, nil

}

func (s *service) GetTicketByDestination(destination string)([]domain.Ticket, error){
	tickets, err := s.repository.GetTicketByDestination(destination)
	if err != nil {
		return []domain.Ticket{}, err
	}
	if len(tickets) == 0{
		return []domain.Ticket{}, fmt.Errorf("empty ticket list for destination: %s", destination)
	}
	
	return tickets, nil
}

func (s *service)GetCountTicketsByDestination(destination string)(int, error){
	countTickets, err := s.repository.GetCountTicketsByDestination(destination)
	if(err != nil){
		return 0, err
	}
	return countTickets, nil
}

func (s *service)GetAVGTicketByDestination(destination string)(float64, error){
	avgTickets, err := s.repository.GetAVGTicketByDestination(destination)
	if(err != nil){
		return 0, err
	}
	return avgTickets, nil
	
}