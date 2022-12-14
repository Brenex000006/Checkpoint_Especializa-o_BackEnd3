package handler

import (
	"errors"
	"strconv"

	"Checkpoint-Backend3/internal/dentista"
	"Checkpoint-Backend3/internal/domain"
	"Checkpoint-Backend3/pkg/web"

	"github.com/gin-gonic/gin"
)

type dentistaHandler struct {
	s dentista.Service
}

func NewDentistaHandler(s dentista.Service) *dentistaHandler {
	return &dentistaHandler{
		s: s,
	}
}

func (h *dentistaHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		dentista, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("dentista not found"))
			return
		}
		web.Success(c, 200, dentista)
	}
}

func validateEmptys2(dentista *domain.Dentista) (bool, error) {
	switch {
	case dentista.Name == "" || dentista.Sobrenome == "" || dentista.Matricula == "":
		return false, errors.New("fields can't be empty")
	}
	return true, nil
}

func (h *dentistaHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dentista domain.Dentista
		err := c.ShouldBindJSON(&dentista)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptys2(&dentista)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		p, err := h.s.Create(dentista)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, p)
	}
}

// Delete elimina un dentista
func (h *dentistaHandler) Delete() gin.HandlerFunc {
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

func (h *dentistaHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		_, err = h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("dentista not found"))
			return
		}
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		var dentista domain.Dentista
		err = c.ShouldBindJSON(&dentista)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptys2(&dentista)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		p, err := h.s.Update(id, dentista)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p)
	}
}

func (h *dentistaHandler) Patch() gin.HandlerFunc {
	type Request struct {
		Name      string `json:"name,omitempty"`
		Sobrenome string `json:"Sobrenome,omitempty"`
		Matricula string `json:"Matricula,omitempty"`
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
			web.Failure(c, 404, errors.New("dentista not found"))
			return
		}
		if err := c.ShouldBindJSON(&r); err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		update := domain.Dentista{
			Name:      r.Name,
			Sobrenome: r.Sobrenome,
			Matricula: r.Matricula,
		}
		p, err := h.s.Update(id, update)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p)
	}
}
