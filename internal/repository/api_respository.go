package repository

import (
	"context"
	"database/sql"
	"microd-api/internal/models"
)

type SQLiteAPIRepository struct {
	db *sql.DB
}

func NewSQLiteAPIRepository(db *sql.DB) APIRepository {
	return &SQLiteAPIRepository{db: db}
}

func (r *SQLiteAPIRepository) CreateAPI(ctx context.Context, api models.API) (int64, error) {
	query := `
		INSERT INTO apis (name, version, description, documentation_link, forum_reference, apm_link, team, tags, swagger)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	result, err := r.db.ExecContext(ctx, query,
		api.Name, api.Version, api.Description, api.DocumentationLink,
		api.ForumReference, api.ApmLink, api.Team, api.Tags, api.Swagger)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (r *SQLiteAPIRepository) GetAPIByID(ctx context.Context, id int64) (models.API, error) {
	query := `SELECT * FROM apis WHERE id = ?`
	var api models.API
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&api.ID, &api.Name, &api.Version, &api.Description, &api.DocumentationLink,
		&api.ForumReference, &api.ApmLink, &api.Team, &api.Tags, &api.Swagger,
		&api.CreatedAt, &api.UpdatedAt)
	return api, err
}

func (r *SQLiteAPIRepository) UpdateAPI(ctx context.Context, api models.API) error {
	query := `
		UPDATE apis
		SET name = ?, version = ?, description = ?, documentation_link = ?,
			forum_reference = ?, apm_link = ?, team = ?, tags = ?, swagger = ?
		WHERE id = ?
	`
	_, err := r.db.ExecContext(ctx, query,
		api.Name, api.Version, api.Description, api.DocumentationLink,
		api.ForumReference, api.ApmLink, api.Team, api.Tags, api.Swagger, api.ID)
	return err
}

func (r *SQLiteAPIRepository) DeleteAPI(ctx context.Context, id int64) error {
	query := `DELETE FROM apis WHERE id = ?`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *SQLiteAPIRepository) ListAPIs(ctx context.Context) ([]models.API, error) {
	query := `SELECT * FROM apis`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var apis []models.API
	for rows.Next() {
		var api models.API
		err := rows.Scan(
			&api.ID, &api.Name, &api.Version, &api.Description, &api.DocumentationLink,
			&api.ForumReference, &api.ApmLink, &api.Team, &api.Tags, &api.Swagger,
			&api.CreatedAt, &api.UpdatedAt)
		if err != nil {
			return nil, err
		}
		apis = append(apis, api)
	}
	return apis, rows.Err()
}
