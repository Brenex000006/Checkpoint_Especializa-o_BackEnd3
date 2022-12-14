package main

import (
	"Checkpoint-Backend3/cmd/server/handler"
	"Checkpoint-Backend3/internal/consulta"
	"Checkpoint-Backend3/internal/dentista"
	"Checkpoint-Backend3/internal/paciente"
	"Checkpoint-Backend3/pkg/consultaPKG"
	"Checkpoint-Backend3/pkg/dentistaPKG"
	"Checkpoint-Backend3/pkg/pacientePKG"

	"github.com/gin-gonic/gin"
)

func main() {

	sqlStorage1 := consultaPKG.NewSQLConsulta()
	sqlStorage2 := pacientePKG.NewSQLPaciente()
	sqlStorage3 := dentistaPKG.NewSQLDentista()

	repo := paciente.NewRepository(sqlStorage2)
	service := paciente.NewService(repo)
	repo2 := dentista.NewRepository(sqlStorage3)
	service2 := dentista.NewService(repo2)
	repo3 := consulta.NewRepository(sqlStorage1)
	service3 := consulta.NewService(repo3)

	consultaHandler := handler.NewConsultaHandler(service3)
	dentistaHandler := handler.NewDentistaHandler(service2)
	pacienteHandler := handler.NewPacienteHandler(service)

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	pacientes := r.Group("/pacientes")
	{
		pacientes.GET(":id", pacienteHandler.GetByID())
		pacientes.POST("", pacienteHandler.Post())
		pacientes.DELETE(":id", pacienteHandler.Delete())
		pacientes.PATCH(":id", pacienteHandler.Patch())
		pacientes.PUT(":id", pacienteHandler.Put())
	}

	dentistas := r.Group("/dentistas")
	{
		dentistas.GET(":id", dentistaHandler.GetByID())
		dentistas.POST("", dentistaHandler.Post())
		dentistas.DELETE(":id", dentistaHandler.Delete())
		dentistas.PATCH(":id", dentistaHandler.Patch())
		dentistas.PUT(":id", dentistaHandler.Put())
	}

	consultas := r.Group("/consultas")
	{
		consultas.GET(":id", consultaHandler.GetByID())
		consultas.POST("", consultaHandler.Post())
		consultas.DELETE(":id", consultaHandler.Delete())
		consultas.PATCH(":id", consultaHandler.Patch())
		consultas.PUT(":id", consultaHandler.Put())
	}

	r.Run(":8080")
}
