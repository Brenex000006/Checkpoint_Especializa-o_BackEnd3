package dentistaPKG

import (
	"encoding/json"
	"errors"
	"os"

	"Checkpoint-Backend3/internal/domain"
)

type jsonDentista struct {
	pathToFile string
}

func (s *jsonDentista) loadDentistas() ([]domain.Dentista, error) {
	var dentistas []domain.Dentista
	file, err := os.ReadFile(s.pathToFile)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(file), &dentistas)
	if err != nil {
		return nil, err
	}
	return dentistas, nil
}

func (s *jsonDentista) saveDentistas(Dentistas []domain.Dentista) error {
	bytes, err := json.Marshal(Dentistas)
	if err != nil {
		return err
	}
	return os.WriteFile(s.pathToFile, bytes, 0644)
}

func NewJsonDentista(path string) DentistaInterface {
	_, err := os.Stat(path)
	if err != nil {
		panic(err)
	}
	return &jsonDentista{
		pathToFile: path,
	}
}

func (s *jsonDentista) Read(id int) (domain.Dentista, error) {
	dentistas, err := s.loadDentistas()
	if err != nil {
		return domain.Dentista{}, err
	}
	for _, Dentista := range dentistas {
		if Dentista.Id == id {
			return Dentista, nil
		}
	}
	return domain.Dentista{}, errors.New("Dentista not found")
}

func (s *jsonDentista) Create(dentista domain.Dentista) error {
	dentistas, err := s.loadDentistas()
	if err != nil {
		return err
	}
	dentista.Id = len(dentistas) + 1
	dentistas = append(dentistas, dentista)
	return s.saveDentistas(dentistas)
}

func (s *jsonDentista) Update(dentista domain.Dentista) error {
	dentistas, err := s.loadDentistas()
	if err != nil {
		return err
	}
	for i, p := range dentistas {
		if p.Id == dentista.Id {
			dentistas[i] = dentista
			return s.saveDentistas(dentistas)
		}
	}
	return errors.New("Dentista not found")
}

func (s *jsonDentista) Delete(id int) error {
	dentistas, err := s.loadDentistas()
	if err != nil {
		return err
	}
	for i, p := range dentistas {
		if p.Id == id {
			dentistas = append(dentistas[:i], dentistas[i+1:]...)
			return s.saveDentistas(dentistas)
		}
	}
	return errors.New("Dentista not found")
}

func (s *jsonDentista) Exists(Matricula string) bool {
	dentistas, err := s.loadDentistas()
	if err != nil {
		return false
	}
	for _, p := range dentistas {
		if p.Matricula == Matricula {
			return true
		}
	}
	return false
}
