package pacientePKG

import "Checkpoint-Backend3/internal/domain"

type PacienteInterface interface {
	Read(id int) (domain.Paciente, error)
	Create(paciente domain.Paciente) error
	Update(paciente domain.Paciente) error
	Delete(id int) error
	Exists(codeValue string) bool
}
