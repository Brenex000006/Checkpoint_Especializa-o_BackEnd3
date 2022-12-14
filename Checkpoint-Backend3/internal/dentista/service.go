package dentista

import (
	"Checkpoint-Backend3/internal/domain"
)

type Service interface {
	GetByID(id int) (domain.Dentista, error)
	Create(p domain.Dentista) (domain.Dentista, error)
	Delete(id int) error
	Update(id int, p domain.Dentista) (domain.Dentista, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetByID(id int) (domain.Dentista, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Dentista{}, err
	}
	return p, nil
}

func (s *service) Create(p domain.Dentista) (domain.Dentista, error) {
	p, err := s.r.Create(p)
	if err != nil {
		return domain.Dentista{}, err
	}
	return p, nil
}

func (s *service) Update(id int, u domain.Dentista) (domain.Dentista, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Dentista{}, err
	}
	if u.Name != "" {
		p.Name = u.Name
	}
	if u.Sobrenome != "" {
		p.Sobrenome = u.Sobrenome
	}
	if u.Matricula != "" {
		p.Matricula = u.Matricula
	}
	p, err = s.r.Update(id, p)
	if err != nil {
		return domain.Dentista{}, err
	}
	return p, nil
}

func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
