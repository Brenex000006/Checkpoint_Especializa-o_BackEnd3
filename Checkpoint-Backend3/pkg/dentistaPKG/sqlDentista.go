package dentistaPKG

import (
	"database/sql"
	"fmt"
	"log"

	"Checkpoint-Backend3/internal/domain"

	_ "github.com/go-sql-driver/mysql"
)

type sqlDentista struct {
	db *sql.DB
	*jsonDentista
}

func NewSQLDentista() DentistaInterface {
	database, err := sql.Open("mysql", "breno:root@tcp(localhost:3306)/clinica")
	if err != nil {
		log.Fatalln(err)
	}

	if err := database.Ping(); err != nil {
		log.Fatalln(err)
	}

	return &sqlDentista{
		db: database,
	}
}

func (s *sqlDentista) Delete(id int) error {
	_, err := s.db.Exec("DELETE FROM Dentistas WHERE id=?", id)
	if err != nil {
		return err
	}

	return nil
}

func (s *sqlDentista) Read(id int) (domain.Dentista, error) {
	dentista := domain.Dentista{}

	rows, err := s.db.Query("SELECT * from Dentistas WHERE id=?", id)
	if err != nil {
		return domain.Dentista{}, err
	}
	for rows.Next() {
		err := rows.Scan(
			&dentista.Id,
			&dentista.Name,
			&dentista.Sobrenome,
			&dentista.Matricula,
		)
		if err != nil {
			return domain.Dentista{}, err
		}
	}
	return dentista, nil
}

func (s *sqlDentista) Update(dentista domain.Dentista) error {
	fmt.Println("updating Dentista")
	_, err := s.db.Exec(
		"UPDATE Dentistas SET name = ?, Sobrenome = ?, Matricula = ? WHERE id = ?;",
		dentista.Name,
		dentista.Sobrenome,
		dentista.Matricula,
		dentista.Id,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (s *sqlDentista) Exists(Matricula string) bool {
	return false
}

func (s *sqlDentista) Create(Dentista domain.Dentista) error {
	_, err := s.db.Exec(
		"INSERT INTO Dentistas (name, Sobrenome, Matricula) VALUES (?, ?, ?)",
		Dentista.Name,
		Dentista.Sobrenome,
		Dentista.Matricula,
	)
	if err != nil {
		return err
	}
	return nil
}
