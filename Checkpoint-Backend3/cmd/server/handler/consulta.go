package handler

import (
	"errors"
	"strconv"
	"strings"

	"Checkpoint-Backend3/internal/consulta"
	"Checkpoint-Backend3/internal/domain"
	"Checkpoint-Backend3/pkg/web"

	"github.com/gin-gonic/gin"
)

type consultaHandler struct {
	s consulta.Service
}

func NewConsultaHandler(s consulta.Service) *consultaHandler {
	return &consultaHandler{
		s: s,
	}
}

func (h *consultaHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		consulta, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("consulta not found"))
			return
		}
		web.Success(c, 200, consulta)
	}
}

func validateEmptys1(consulta *domain.Consulta) (bool, error) {
	switch {
	case consulta.Paciente == "" || consulta.Dentista == "" || consulta.Descricao == "" || consulta.DataHora == "":
		return false, errors.New("fields can't be empty")
	}
	return true, nil
}

func validateDataHora(exp string) (bool, error) {
	dates := strings.Split(exp, "/")
	list := []int{}
	if len(dates) != 3 {
		return false, errors.New("invalid date, must be in format: dd/mm/yyyy")
	}
	for value := range dates {
		number, err := strconv.Atoi(dates[value])
		if err != nil {
			return false, errors.New("invalid date, must be numbers")
		}
		list = append(list, number)
	}
	condition := (list[0] < 1 || list[0] > 31) && (list[1] < 1 || list[1] > 12) && (list[2] < 1 || list[2] > 9999)
	if condition {
		return false, errors.New("invalid date, date must be between 1 and 31/12/9999")
	}
	return true, nil
}

func (h *consultaHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var consulta domain.Consulta
		err := c.ShouldBindJSON(&consulta)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptys1(&consulta)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		valid, err = validateDataHora(consulta.DataHora)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		p, err := h.s.Create(consulta)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, p)
	}
}

func (h *consultaHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		err = h.s.Delete(id)
		if err != nil {
			web.Failure(c, 404, err)
			return
		}
		web.Success(c, 204, nil)
	}
}

func (h *consultaHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		_, err = h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("consulta not found"))
			return
		}
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		var consulta domain.Consulta
		err = c.ShouldBindJSON(&consulta)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptys1(&consulta)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		valid, err = validateDataHora(consulta.DataHora)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		p, err := h.s.Update(id, consulta)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p)
	}
}

func (h *consultaHandler) Patch() gin.HandlerFunc {
	type Request struct {
		DataHora  string `json:"DataHora,omitempty"`
		Paciente  string `json:"Paciente,omitempty"`
		Descricao string `json:"Descricao,omitempty"`
		Dentista  string `json:"Dentista,omitempty"`
	}
	return func(c *gin.Context) {
		var r Request
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		_, err = h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("consulta not found"))
			return
		}
		if err := c.ShouldBindJSON(&r); err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		update := domain.Consulta{
			DataHora:  r.DataHora,
			Paciente:  r.Paciente,
			Descricao: r.Descricao,
			Dentista:  r.Dentista,
		}
		if update.DataHora != "" {
			valid, err := validateDataHora(update.DataHora)
			if !valid {
				web.Failure(c, 400, err)
				return
			}
		}
		p, err := h.s.Update(id, update)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p)
	}
}
