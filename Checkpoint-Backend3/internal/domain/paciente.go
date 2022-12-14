package domain

type Paciente struct {
	Id             int    `json:"id"`
	Name           string `json:"name" binding:"required"`
	Sobrenome      string `json:"sobrenome" binding:"required"`
	RG             string `json:"rg" binding:"required"`
	DatadeCadastro string `json:"datadecadastro" binding:"required"`
}
