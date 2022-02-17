package rest

import (
	"fmt"
	"log"

	_ "github.com/c-4u/pinned-attendant/app/rest/docs"
	"github.com/c-4u/pinned-attendant/domain/service"
	"github.com/c-4u/pinned-attendant/infra/client/kafka"
	"github.com/c-4u/pinned-attendant/infra/db"
	"github.com/c-4u/pinned-attendant/infra/repo"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// @title Attendant Swagger API
// @version 1.0
// @description Swagger API for Attendant Service.
// @termsOfService http://swagger.io/terms/

// @contact.name Coding4u
// @contact.email contato@coding4u.com.br

// @BasePath /api/v1
// @query.collection.format multi

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func StartRestServer(pg *db.PostgreSQL, kp *kafka.KafkaProducer, port int) {
	r := fiber.New()
	r.Use(cors.New())

	repository := repo.NewRepository(pg, kp)
	service := service.NewService(repository)
	restService := NewRestService(service)

	api := r.Group("/api")

	v1 := api.Group("/v1")
	v1.Get("/swagger/*", fiberSwagger.WrapHandler)
	{
		attendant := v1.Group("/attendants")
		attendant.Post("", restService.CreateAttendant)
		attendant.Get("/:attendant_id", restService.FindAttendant)
	}

	addr := fmt.Sprintf("0.0.0.0:%d", port)
	err := r.Listen(addr)
	if err != nil {
		log.Fatal("cannot start rest server", err)
	}

	log.Printf("rest server has been started on port %d", port)
}
