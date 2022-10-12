package main

import (
	"log"
	"os"

	"github.com/ZoeTira/desafio-goweb-ZoeTira/cmd/server/handler"
	"github.com/ZoeTira/desafio-goweb-ZoeTira/docs"
	"github.com/ZoeTira/desafio-goweb-ZoeTira/internal/tickets"
	"github.com/ZoeTira/desafio-goweb-ZoeTira/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/joho/godotenv"
)

// @title Challenge GoWeb API
// @version 1.0
// @description This API handles tickets.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

func main(){
	err := godotenv.Load()
	if(err != nil){
		log.Fatal(err.Error())
	}
	bd := store.New(store.FileType, "./tickets.csv")

	repository := tickets.NewRepository(bd)
	service := tickets.NewService(repository)
	controller := handler.NewTicket(service)

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	r := gin.Default()
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	pr := r.Group("/tickets")
	{
		pr.GET("/list", controller.GetAll())
		pr.GET("/getByCountry/:country", controller.GetTicketByDestination())
		pr.GET("/getCountByCountry/:country", controller.GetCountTicketsByDestination())
		pr.GET("/getAVGByCountry/:country", controller.GetAVGTicketByDestination())
	}
	r.Run()

}