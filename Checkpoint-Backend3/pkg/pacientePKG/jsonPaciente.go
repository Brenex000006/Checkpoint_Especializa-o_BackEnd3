package pacientePKG

import (
	"encoding/json"
	"errors"
	"os"

	"Checkpoint-Backend3/internal/domain"
)

type jsonPaciente struct {
	pathToFile string
}

func (s *jsonPaciente) loadPacientes() ([]domain.Paciente, error) {
	var pacientes []domain.Paciente
	file, err := os.ReadFile(s.pathToFile)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(file), &pacientes)
	if err != nil {
		return nil, err
	}
	return pacientes, nil
}

func (s *jsonPaciente) savePacientes(Pacientes []domain.Paciente) error {
	bytes, err := json.Marshal(Pacientes)
	if err != nil {
		return err
	}
	return os.WriteFile(s.pathToFile, bytes, 0644)
}

func NewJsonPaciente(path string) PacienteInterface {
	_, err := os.Stat(path)
	if err != nil {
		panic(err)
	}
	return &jsonPaciente{
		pathToFile: path,
	}
}

func (s *jsonPaciente) Read(id int) (domain.Paciente, error) {
	pacientes, err := s.loadPacientes()
	if err != nil {
		return domain.Paciente{}, err
	}
	for _, paciente := range pacientes {
		if paciente.Id == id {
			return paciente, nil
		}
	}
	return domain.Paciente{}, errors.New("Paciente not found")
}

func (s *jsonPaciente) Create(paciente domain.Paciente) error {
	pacientes, err := s.loadPacientes()
	if err != nil {
		return err
	}
	paciente.Id = len(pacientes) + 1
	pacientes = append(pacientes, paciente)
	return s.savePacientes(pacientes)
}

func (s *jsonPaciente) Update(paciente domain.Paciente) error {
	pacientes, err := s.loadPacientes()
	if err != nil {
		return err
	}
	for i, p := range pacientes {
		if p.Id == paciente.Id {
			pacientes[i] = paciente
			return s.savePacientes(pacientes)
		}
	}
	return errors.New("Paciente not found")
}

func (s *jsonPaciente) Delete(id int) error {
	pacientes, err := s.loadPacientes()
	if err != nil {
		return err
	}
	for i, p := range pacientes {
		if p.Id == id {
			pacientes = append(pacientes[:i], pacientes[i+1:]...)
			return s.savePacientes(pacientes)
		}
	}
	return errors.New("Paciente not found")
}

func (s *jsonPaciente) Exists(RG string) bool {
	pacientes, err := s.loadPacientes()
	if err != nil {
		return false
	}
	for _, p := range pacientes {
		if p.RG == RG {
			return true
		}
	}
	return false
}
