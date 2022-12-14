package paciente

import (
	"errors"
	"fmt"

	"Checkpoint-Backend3/internal/domain"
	"Checkpoint-Backend3/pkg/pacientePKG"
)

type Repository interface {
	GetByID(id int) (domain.Paciente, error)
	Create(p domain.Paciente) (domain.Paciente, error)
	Update(id int, p domain.Paciente) (domain.Paciente, error)
	Delete(id int) error
}

type repository struct {
	storage pacientePKG.PacienteInterface
}

func NewRepository(storage pacientePKG.PacienteInterface) Repository {
	return &repository{storage}
}

func (r *repository) GetByID(id int) (domain.Paciente, error) {
	paciente, err := r.storage.Read(id)
	if err != nil {
		return domain.Paciente{}, errors.New("Paciente not found")
	}
	return paciente, nil

}

func (r *repository) Create(p domain.Paciente) (domain.Paciente, error) {
	err := r.storage.Create(p)
	if err != nil {
		return domain.Paciente{}, errors.New(fmt.Sprintf("error creating Paciente: %s", err.Error()))
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

func (r *repository) Update(id int, p domain.Paciente) (domain.Paciente, error) {
	err := r.storage.Update(p)
	if err != nil {
		return domain.Paciente{}, errors.New("error .. updating Paciente")
	}
	return p, nil
}
