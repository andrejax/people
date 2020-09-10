package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"people/interfaces"
	"people/models"
	er "people/utils/errors"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(Conn *sql.DB) interfaces.UserRepository {
	return &UserRepository{Conn}
}

func (m *UserRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []models.User, err error) {
	rows, err := m.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer func() {
		rows.Close()
	}()

	result = make([]models.User, 0)
	for rows.Next() {
		u := models.User{}
		err = rows.Scan(
			&u.ID,
			&u.Name,
			&u.Email,
			&u.Password,
		)

		if err != nil {
			return nil, err
		}

		result = append(result, u)
	}

	return result, nil
}


func (ur *UserRepository) Get(ctx context.Context, id int64)  (models.User, error) {
	query := `SELECT id, name, email, password FROM user WHERE id = ?`

	list, err := ur.fetch(ctx, query, id)
	if err != nil {
		return models.User{}, err
	}

	if len(list) > 0 {
		return list[0], nil
	}

	return models.User{}, er.NotFound
}

func (ur *UserRepository) List(ctx context.Context)  (res []models.User, err error) {
	query := `SELECT id, name, email, password FROM user `
	res, err = ur.fetch(ctx, query)
	if err != nil {
		return nil, err
	}

	return
}

func (ur *UserRepository) Add(ctx context.Context, u *models.User) (err error) {
	query := `INSERT user SET name=? , email=? , password=?`
	stmt, err := ur.DB.PrepareContext(ctx, query)

	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, u.Name, u.Email, u.Password)
	if err != nil {
		return
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return
	}

	u.ID = lastID
	return
}

func (ur *UserRepository) Update(ctx context.Context, u *models.User) (err error) {
	query := `UPDATE user set name=?, email=?, password=? WHERE ID = ?`

	stmt, err := ur.DB.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, u.Name, u.Email, u.Password, u.ID)
	if err != nil {
		return
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return
	}

	if affect != 1 {
		err = fmt.Errorf("Something went wrong. Total Affected: %d", affect)
		return
	}

	return
}

func (ur *UserRepository) Remove(ctx context.Context, id int64) (err error) {
	query := "DELETE FROM user WHERE id = ?"

	stmt, err := ur.DB.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return
	}

	rowsAfected, err := res.RowsAffected()
	if err != nil {
		return
	}

	if rowsAfected != 1 {
		err = fmt.Errorf("Something went wrong. Total Affected: %d", rowsAfected)
		return
	}

	return
}