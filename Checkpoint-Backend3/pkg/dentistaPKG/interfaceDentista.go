package dentistaPKG

import "Checkpoint-Backend3/internal/domain"

type DentistaInterface interface {
	Read(id int) (domain.Dentista, error)
	Create(dentista domain.Dentista) error
	Update(dentista domain.Dentista) error
	Delete(id int) error
	Exists(codeValue string) bool
}
