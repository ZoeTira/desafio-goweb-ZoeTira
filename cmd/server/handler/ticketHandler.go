package handler

import (
	"net/http"
	"os"

	//"strconv"

	//"github.com/ZoeTira/desafio-goweb-ZoeTira/domain"
	"github.com/ZoeTira/desafio-goweb-ZoeTira/domain"
	"github.com/ZoeTira/desafio-goweb-ZoeTira/internal/tickets"
	"github.com/gin-gonic/gin"
)

type Ticket struct{
	service tickets.Service
}

func NewTicket(p tickets.Service) *Ticket{
	return &Ticket{
		service: p,
	}
	
}

func(ticket *Ticket) GetAll() gin.HandlerFunc{
	return func(ctx *gin.Context){
		token := ctx.GetHeader("token")
		if(token != os.Getenv("TOKEN")){
			ctx.JSON(http.StatusNonAuthoritativeInfo, "Invalid Token")
		}
		
		t, err := ticket.service.GetAll()
		if(err != nil){
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return 
		}

		if (len(t) == 0){
			ctx.JSON(http.StatusNotFound, "empty ticket list")
		}

		ctx.JSON(http.StatusOK, t)
	}
}

func (ticket *Ticket) GetTicketByDestination() gin.HandlerFunc{
	return func(ctx *gin.Context){
		token := ctx.GetHeader("token")
		if(token != os.Getenv("TOKEN")){
			ctx.JSON(http.StatusNonAuthoritativeInfo, "Invalid Token")
		}

		var req domain.Ticket
		t, err := ticket.service.GetTicketByDestination(req.Country)
		if(err != nil){
			ctx.JSON(http.StatusInternalServerError, err.Error())
		}

		if (len(t) == 0){
			ctx.JSON(http.StatusNotFound, "empty ticket list for destination")
		}
		 
		ctx.JSON(http.StatusOK, t)

	}
}