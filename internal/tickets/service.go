package tickets

import (
	"fmt"

	"github.com/ZoeTira/desafio-goweb-ZoeTira/domain"
)

type Service interface{
	GetAll()([]domain.Ticket, error)
	GetTicketByDestination(destination string)([]domain.Ticket, error)
	
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
	t,err := s.repository.GetAll()
	if(err != nil){
		return []domain.Ticket{}, err
	}

	if len(t) == 0{
		return []domain.Ticket{}, fmt.Errorf("empty ticket list")
	}

	return t, nil

}

func (s *service) GetTicketByDestination(destination string)([]domain.Ticket, error){
	t, err := s.repository.GetTicketByDestination(destination)
	if err != nil {
		return []domain.Ticket{}, err
	}
	if len(t) == 0{
		return []domain.Ticket{}, fmt.Errorf("empty ticket list for destination: %s", destination)
	}
	
	return t, nil
}