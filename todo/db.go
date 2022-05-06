package todoapi

import (
	"database/sql"
	todo "todo/gen/todo"
)

type Sql struct {
	db *sql.DB
}

func NewSqlDb(db *sql.DB) *Sql {
	return &Sql{db}
}

func (s *Sql) Find(id int) (*todo.Username, error) {
	var t todo.Username
	err := s.db.QueryRow("select id ,name from users where id = ?", id).Scan(&t.ID, &t.Name)
	if err != nil {
		return nil, err
	}
	return &t, nil

}
