package consulta

import (
	"Checkpoint-Backend3/internal/domain"
)

type Service interface {
	GetByID(id int) (domain.Consulta, error)
	Create(p domain.Consulta) (domain.Consulta, error)
	Delete(id int) error
	Update(id int, p domain.Consulta) (domain.Consulta, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetByID(id int) (domain.Consulta, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Consulta{}, err
	}
	return p, nil
}

func (s *service) Create(p domain.Consulta) (domain.Consulta, error) {
	p, err := s.r.Create(p)
	if err != nil {
		return domain.Consulta{}, err
	}
	return p, nil
}
func (s *service) Update(id int, u domain.Consulta) (domain.Consulta, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Consulta{}, err
	}
	if u.DataHora != "" {
		p.DataHora = u.DataHora
	}
	if u.Paciente != "" {
		p.Paciente = u.Paciente
	}
	if u.Descricao != "" {
		p.Descricao = u.Descricao
	}
	if u.Dentista != "" {
		p.Dentista = u.Dentista
	}
	p, err = s.r.Update(id, p)
	if err != nil {
		return domain.Consulta{}, err
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
