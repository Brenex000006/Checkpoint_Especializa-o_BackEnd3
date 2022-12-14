package consulta

import (
	"errors"
	"fmt"

	"Checkpoint-Backend3/internal/domain"
	"Checkpoint-Backend3/pkg/consultaPKG"
)

type Repository interface {
	GetByID(id int) (domain.Consulta, error)
	Create(p domain.Consulta) (domain.Consulta, error)
	Update(id int, p domain.Consulta) (domain.Consulta, error)
	Delete(id int) error
}

type repository struct {
	storage consultaPKG.ConsultaInterface
}

func NewRepository(storage consultaPKG.ConsultaInterface) Repository {
	return &repository{storage}
}

func (r *repository) GetByID(id int) (domain.Consulta, error) {
	consulta, err := r.storage.Read(id)
	if err != nil {
		return domain.Consulta{}, errors.New("Consulta not found")
	}
	return consulta, nil

}

func (r *repository) Create(p domain.Consulta) (domain.Consulta, error) {
	err := r.storage.Create(p)
	if err != nil {
		return domain.Consulta{}, errors.New(fmt.Sprintf("error creating Consulta: %s", err.Error()))
	}
	return p, nil
}

func (r *repository) Delete(id int) error {
	err := r.storage.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Update(id int, p domain.Consulta) (domain.Consulta, error) {
	err := r.storage.Update(p)
	if err != nil {
		return domain.Consulta{}, errors.New("error updating Consulta")
	}
	return p, nil
}
