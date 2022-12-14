package domain

type Dentista struct {
	Id        int    `json:"id"`
	Name      string `json:"name" binding:"required"`
	Sobrenome string `json:"sobrenome" binding:"required"`
	Matricula string `json:"matricula" binding:"required"`
}
