package consultaPKG

import (
	"database/sql"
	"fmt"
	"log"

	"Checkpoint-Backend3/internal/domain"

	_ "github.com/go-sql-driver/mysql"
)

type sqlConsulta struct {
	db *sql.DB
	*jsonConsulta
}

func NewSQLConsulta() ConsultaInterface {
	database, err := sql.Open("mysql", "breno:root@tcp(localhost:3306)/clinica")
	if err != nil {
		log.Fatalln(err)
	}

	if err := database.Ping(); err != nil {
		log.Fatalln(err)
	}

	return &sqlConsulta{
		db: database,
	}
}

func (s *sqlConsulta) Delete(id int) error {
	_, err := s.db.Exec("DELETE FROM Consultas WHERE id=?", id)
	if err != nil {
		return err
	}

	return nil
}

func (s *sqlConsulta) Read(id int) (domain.Consulta, error) {
	consulta := domain.Consulta{}

	rows, err := s.db.Query("SELECT * from Consultas WHERE id=?", id)
	if err != nil {
		return domain.Consulta{}, err
	}
	for rows.Next() {
		err := rows.Scan(
			&consulta.Id,
			&consulta.DataHora,
			&consulta.Paciente,
			&consulta.Descricao,
			&consulta.Dentista,
		)
		if err != nil {
			return domain.Consulta{}, err
		}
	}
	return consulta, nil
}

func (s *sqlConsulta) Update(consulta domain.Consulta) error {
	fmt.Println("updating Consulta")
	_, err := s.db.Exec(
		"UPDATE Consultas SET Paciente = ?, DataHora = ?, Descricao = ?, Dentista = ? WHERE id = ?;",
		consulta.Paciente,
		consulta.DataHora,
		consulta.Descricao,
		consulta.Dentista,
		consulta.Id,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (s *sqlConsulta) Exists(Descricao string) bool {
	return false
}

func (s *sqlConsulta) Create(consulta domain.Consulta) error {
	_, err := s.db.Exec(
		"INSERT INTO Consultas (Paciente, DataHora, Descricao, Dentista) VALUES (?, ?, ?, ?)",
		consulta.Paciente,
		consulta.DataHora,
		consulta.Descricao,
		consulta.Dentista,
	)
	if err != nil {
		return err
	}
	return nil
}
