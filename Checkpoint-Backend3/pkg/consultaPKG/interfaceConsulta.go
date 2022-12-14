package consultaPKG

import "Checkpoint-Backend3/internal/domain"

type ConsultaInterface interface {
	Read(id int) (domain.Consulta, error)
	Create(consulta domain.Consulta) error
	Update(consulta domain.Consulta) error
	Delete(id int) error
	Exists(codeValue string) bool
}
