package consultaPKG

import (
	"encoding/json"
	"errors"
	"os"

	"Checkpoint-Backend3/internal/domain"
)

type jsonConsulta struct {
	pathToFile string
}

func (s *jsonConsulta) loadConsultas() ([]domain.Consulta, error) {
	var consultas []domain.Consulta
	file, err := os.ReadFile(s.pathToFile)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(file), &consultas)
	if err != nil {
		return nil, err
	}
	return consultas, nil
}

func (s *jsonConsulta) saveConsultas(Consultas []domain.Consulta) error {
	bytes, err := json.Marshal(Consultas)
	if err != nil {
		return err
	}
	return os.WriteFile(s.pathToFile, bytes, 0644)
}

func NewJsonConsulta(path string) ConsultaInterface {
	_, err := os.Stat(path)
	if err != nil {
		panic(err)
	}
	return &jsonConsulta{
		pathToFile: path,
	}
}

func (s *jsonConsulta) Read(id int) (domain.Consulta, error) {
	consultas, err := s.loadConsultas()
	if err != nil {
		return domain.Consulta{}, err
	}
	for _, consulta := range consultas {
		if consulta.Id == id {
			return consulta, nil
		}
	}
	return domain.Consulta{}, errors.New("Consulta not found")
}

func (s *jsonConsulta) Create(consulta domain.Consulta) error {
	consultas, err := s.loadConsultas()
	if err != nil {
		return err
	}
	consulta.Id = len(consultas) + 1
	consultas = append(consultas, consulta)
	return s.saveConsultas(consultas)
}

func (s *jsonConsulta) Update(consulta domain.Consulta) error {
	consultas, err := s.loadConsultas()
	if err != nil {
		return err
	}
	for i, p := range consultas {
		if p.Id == consulta.Id {
			consultas[i] = consulta
			return s.saveConsultas(consultas)
		}
	}
	return errors.New("Consulta not found")
}

func (s *jsonConsulta) Delete(id int) error {
	consultas, err := s.loadConsultas()
	if err != nil {
		return err
	}
	for i, p := range consultas {
		if p.Id == id {
			consultas = append(consultas[:i], consultas[i+1:]...)
			return s.saveConsultas(consultas)
		}
	}
	return errors.New("Consulta not found")
}

func (s *jsonConsulta) Exists(Descricao string) bool {
	consultas, err := s.loadConsultas()
	if err != nil {
		return false
	}
	for _, p := range consultas {
		if p.Descricao == Descricao {
			return true
		}
	}
	return false
}
