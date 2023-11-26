package sql

import (
	"database/sql"
	"errors"
	"server/api/models"
)

type UserModel struct {
	Db *sql.DB
}

func (m *UserModel) Insert(name, surname string, age int) (int, error) {
	query := `INSERT INTO users(name, surname, age) VALUES(?, ?, ?)`
	result, err := m.Db.Exec(query, name, surname, age)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (m *UserModel) Get(id int) (*models.User, error) {
	query := `SELECT id, name, surname, age FROM users WHERE id = ?`
	user := models.User{}
	if err := m.Db.QueryRow(query, id).Scan(&user.Id, &user.Name, &user.Surname, &user.Age); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		}
		return nil, err
	}
	return &user, nil
}

func (m *UserModel) Latest() ([]*models.User, error) {
	query := `SELECT id, name, surname, age FROM users ORDER BY age DESC LIMIT 10`
	rows, err := m.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var snippets []*models.User
	for rows.Next() {
		s := &models.User{}
		err = rows.Scan(&s.Id, &s.Name, &s.Surname, &s.Age)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, s)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return snippets, nil
}
