package dentista

import (
	"errors"
	"fmt"

	"Checkpoint-Backend3/internal/domain"
	"Checkpoint-Backend3/pkg/dentistaPKG"
)

type Repository interface {
	GetByID(id int) (domain.Dentista, error)
	Create(p domain.Dentista) (domain.Dentista, error)
	Update(id int, p domain.Dentista) (domain.Dentista, error)
	Delete(id int) error
}

type repository struct {
	storage dentistaPKG.DentistaInterface
}

func NewRepository(storage dentistaPKG.DentistaInterface) Repository {
	return &repository{storage}
}

func (r *repository) GetByID(id int) (domain.Dentista, error) {
	dentista, err := r.storage.Read(id)
	if err != nil {
		return domain.Dentista{}, errors.New("Dentista not found")
	}
	return dentista, nil

}

func (r *repository) Create(p domain.Dentista) (domain.Dentista, error) {
	err := r.storage.Create(p)
	if err != nil {
		return domain.Dentista{}, errors.New(fmt.Sprintf("error creating Dentista: %s", err.Error()))
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

func (r *repository) Update(id int, p domain.Dentista) (domain.Dentista, error) {
	err := r.storage.Update(p)
	if err != nil {
		return domain.Dentista{}, errors.New("error updating Dentista")
	}
	return p, nil
}
