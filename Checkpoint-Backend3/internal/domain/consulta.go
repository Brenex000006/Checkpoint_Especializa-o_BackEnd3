package domain

type Consulta struct {
	Id        int    `json:"id"`
	DataHora  string `json:"DataHora" binding:"required"`
	Paciente  string `json:"Paciente" binding:"required"`
	Descricao string `json:"Descricao" binding:"required"`
	Dentista  string `json:"Dentista" binding:"required"`
}
