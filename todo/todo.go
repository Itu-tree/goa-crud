package todoapi

import (
	"context"
	"database/sql"
	"log"
	todo "todo/gen/todo"
)

// todo service example implementation.
// The example methods log the requests and return zero values.
type todosrvc struct {
	db     *Sql
	logger *log.Logger
}

// NewTodo returns the todo service implementation.
func NewTodo(db *sql.DB, logger *log.Logger) todo.Service {
	sql := NewSqlDb(db)
	return &todosrvc{sql, logger}
}

// Hello implements hello.
func (s *todosrvc) Hello(ctx context.Context, p *todo.HelloPayload) (res string, err error) {
	s.logger.Print("todo.hello")
	return
}

// Show implements show.
func (s *todosrvc) Show(ctx context.Context, p *todo.ShowPayload) (res *todo.Username, err error) {
	s.logger.Print("todo.show")
	t, err := s.db.Find(p.ID)
	if err != nil {
		return nil, err
	}
	return t, nil
}
