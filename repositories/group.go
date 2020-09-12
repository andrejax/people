package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"people/interfaces"
	"people/models"
	er "people/utils/errors"
)

type GroupRepository struct {
	DB *sql.DB
}

func NewGroupRepository(Conn *sql.DB) interfaces.GroupRepository {
	return &GroupRepository{Conn}
}

func (gr *GroupRepository) fetch(ctx context.Context, query string, args ...interface{}) (result []models.Group, err error) {
	rows, err := gr.DB.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer func() {
		rows.Close()
	}()

	result = make([]models.Group, 0)
	for rows.Next() {
		u := models.Group{}
		err = rows.Scan(
			&u.Id,
			&u.Name,
		)

		if err != nil {
			return nil, err
		}

		result = append(result, u)
	}

	return result, nil
}


func (gr *GroupRepository) Get(ctx context.Context, id string)  (models.Group, error) {
	query := `SELECT id, name FROM users_group WHERE id = ?`

	list, err := gr.fetch(ctx, query, id)
	if err != nil {
		return models.Group{}, err
	}

	if len(list) > 0 {
		return list[0], nil
	}

	return models.Group{}, er.NotFound
}

func (gr *GroupRepository) List(ctx context.Context)  (res []models.Group, err error) {
	query := `SELECT * FROM users_group`
	res, err = gr.fetch(ctx, query)
	if err != nil {
		return nil, err
	}

	return
}

func (gr *GroupRepository) Add(ctx context.Context, g *models.Group) (err error) {
	query := `INSERT users_group SET name=? `
	stmt, err := gr.DB.PrepareContext(ctx, query)

	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, g.Name)
	if err != nil {
		return
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return
	}

	g.Id = string(lastID)
	return
}

func (gr *GroupRepository) Update(ctx context.Context, g *models.Group) (err error) {
	query := `UPDATE users_group set name=? WHERE Id = ?`

	stmt, err := gr.DB.PrepareContext(ctx, query)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, g.Name, g.Id)
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

func (gr *GroupRepository) Remove(ctx context.Context, id string) (err error) {
	query := "DELETE FROM users_group WHERE id = ?"

	stmt, err := gr.DB.PrepareContext(ctx, query)
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