package handler

import (
	//"fmt"
	"net/http"
	"os"

	//"strconv"

	//"github.com/ZoeTira/desafio-goweb-ZoeTira/domain"
	//"github.com/ZoeTira/desafio-goweb-ZoeTira/domain"
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
// ListTickets godoc
// @Summary List tickets
// @Tags Tickets
// @Description get tickets
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} []domain.Ticket
// @Router /tickets/list [get]
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

// ListTicketsByDestination godoc
// @Summary List tickets by destination
// @Tags Tickets
// @Description get tickets by destination
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} []domain.Ticket
// @Router /tickets/getByCountry/:country [get]
func (ticket *Ticket) GetTicketByDestination() gin.HandlerFunc{
	return func(ctx *gin.Context){
		token := ctx.GetHeader("token")
		if(token != os.Getenv("TOKEN")){
			ctx.JSON(http.StatusNonAuthoritativeInfo, "Invalid Token")
		}

		//var req domain.Ticket
		country := ctx.Param("country")
		t, err := ticket.service.GetTicketByDestination(country)
		if(err != nil){
			ctx.JSON(http.StatusInternalServerError, err.Error())
		}

		if (len(t) == 0){
			ctx.JSON(http.StatusNotFound, "empty ticket list for destination")
		}
		 
		ctx.JSON(http.StatusOK, t)

	}
}

// CountTicketsByDestination godoc
// @Summary Count tickets by destination
// @Tags Tickets
// @Description get count of tickets by destination
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} []domain.Ticket
// @Router /tickets/getCountByCountry/:country [get]
func (ticket *Ticket) GetCountTicketsByDestination() gin.HandlerFunc {
	return func(ctx *gin.Context){
		token := ctx.GetHeader("token")
		if(token != os.Getenv("TOKEN")){
			ctx.JSON(http.StatusNonAuthoritativeInfo, "Invalid Token")
		}

		country := ctx.Param("country")
		countTickets, err := ticket.service.GetCountTicketsByDestination(country)
		if(err != nil){
			ctx.JSON(http.StatusInternalServerError, err.Error())
		}

		ctx.JSON(http.StatusOK, countTickets)

	}
	
}

// AVGTicketsByDestination godoc
// @Summary AVG tickets by destination
// @Tags Tickets
// @Description get average number of tickets per destination compared to all tickets
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} []domain.Ticket
// @Router /tickets/getAVGByCountry/:country [get]
func (ticket *Ticket)GetAVGTicketByDestination() gin.HandlerFunc{
	return func(ctx *gin.Context){
		token := ctx.GetHeader("token")
		if(token != os.Getenv("TOKEN")){
			ctx.JSON(http.StatusNonAuthoritativeInfo, "Invalid Token")
		}

		country := ctx.Param("country")
		avgTickets, err := ticket.service.GetAVGTicketByDestination(country)
		if(err != nil){
			ctx.JSON(http.StatusInternalServerError, err.Error())
		}
	//fmt.Println(avgTickets)
		ctx.JSON(http.StatusOK, avgTickets)

	}
}