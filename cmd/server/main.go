package main

import (
	"log"

	"github.com/ZoeTira/desafio-goweb-ZoeTira/cmd/server/handler"
	"github.com/ZoeTira/desafio-goweb-ZoeTira/internal/tickets"
	"github.com/ZoeTira/desafio-goweb-ZoeTira/pkg/store"
	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
)

func main(){
	err := godotenv.Load()
	if(err != nil){
		log.Fatal(err.Error())
	}
	bd := store.New(store.FileType, "./tickets.csv")

	repository := tickets.NewRepository(bd)
	service := tickets.NewService(repository)
	controller := handler.NewTicket(service)

	r := gin.Default()
	pr := r.Group("/tickets")
	{
		pr.GET("/list", controller.GetAll())
		pr.GET("/getByCountry/:country", controller.GetTicketByDestination())
		pr.GET("/getCountByCountry/:country", controller.GetCountTicketsByDestination())
		pr.GET("/getAVGByCountry/:country", controller.GetAVGTicketByDestination())
	}
	r.Run()

}