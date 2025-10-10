package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"task-management/helper"
	"task-management/model/domain"
)

type ProjectRepositoryImpl struct {
	DB *sql.DB
}

func NewProjectRepository(db *sql.DB) ProjectRepository {
	return &ProjectRepositoryImpl{DB: db}
}

func (r *ProjectRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, project domain.Project) domain.Project {
	if project.Id == uuid.Nil {
		project.Id = uuid.New()
	}
	project.CreatedAt = time.Now()
	project.UpdatedAt = time.Now()

	SQL := `INSERT INTO projects(
		id, name, description, progress, confidence, trend, created_at, updated_at, user_id
	) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	var err error
	if tx != nil {
		_, err = tx.ExecContext(ctx, SQL,
			project.Id,
			project.Name,
			project.Description,
			project.Progress,
			project.Confidence,
			project.Trend,
			project.CreatedAt,
			project.UpdatedAt,
			project.UserId,
		)
	} else {
		_, err = r.DB.ExecContext(ctx, SQL,
			project.Id,
			project.Name,
			project.Description,
			project.Progress,
			project.Confidence,
			project.Trend,
			project.CreatedAt,
			project.UpdatedAt,
			project.UserId,
		)
	}
	helper.PanicIfError(err)
	return project
}

func (r *ProjectRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, project domain.Project) domain.Project {
	project.UpdatedAt = time.Now()

	SQL := `UPDATE projects 
			SET name = $1, description = $2, progress = $3, confidence = $4, trend = $5, updated_at = $6 
			WHERE id = $7`

	var err error
	if tx != nil {
		_, err = tx.ExecContext(ctx, SQL,
			project.Name,
			project.Description,
			project.Progress,
			project.Confidence,
			project.Trend,
			project.UpdatedAt,
			project.Id,
		)
	} else {
		_, err = r.DB.ExecContext(ctx, SQL,
			project.Name,
			project.Description,
			project.Progress,
			project.Confidence,
			project.Trend,
			project.UpdatedAt,
			project.Id,
		)
	}
	helper.PanicIfError(err)
	return project
}

func (r *ProjectRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, projectId uuid.UUID) error {
	SQL := "DELETE FROM projects WHERE id = $1"
	var err error
	if tx != nil {
		_, err = tx.ExecContext(ctx, SQL, projectId)
	} else {
		_, err = r.DB.ExecContext(ctx, SQL, projectId)
	}
	if err != nil {
		return err
	}
	return nil
}

func (r *ProjectRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, projectId uuid.UUID) (domain.Project, error) {
	SQL := `SELECT id, name, description, progress, confidence, trend, created_at, updated_at, user_id 
			FROM projects WHERE id = $1`

	var project domain.Project
	var err error
	if tx != nil {
		err = tx.QueryRowContext(ctx, SQL, projectId).Scan(
			&project.Id,
			&project.Name,
			&project.Description,
			&project.Progress,
			&project.Confidence,
			&project.Trend,
			&project.CreatedAt,
			&project.UpdatedAt,
			&project.UserId,
		)
	} else {
		err = r.DB.QueryRowContext(ctx, SQL, projectId).Scan(
			&project.Id,
			&project.Name,
			&project.Description,
			&project.Progress,
			&project.Confidence,
			&project.Trend,
			&project.CreatedAt,
			&project.UpdatedAt,
			&project.UserId,
		)
	}

	if err == sql.ErrNoRows {
		return project, errors.New("project not found")
	}
	helper.PanicIfError(err)
	return project, nil
}

func (r *ProjectRepositoryImpl) FindByUserId(ctx context.Context, tx *sql.Tx, userId uuid.UUID) []domain.Project {
	SQL := `SELECT id, name, description, progress, confidence, trend, created_at, updated_at, user_id 
			FROM projects WHERE user_id = $1`

	var projects []domain.Project
	var rows *sql.Rows
	var err error

	if tx != nil {
		rows, err = tx.QueryContext(ctx, SQL, userId)
	} else {
		rows, err = r.DB.QueryContext(ctx, SQL, userId)
	}

	helper.PanicIfError(err)
	defer rows.Close()

	for rows.Next() {
		var project domain.Project
		err := rows.Scan(
			&project.Id,
			&project.Name,
			&project.Description,
			&project.Progress,
			&project.Confidence,
			&project.Trend,
			&project.CreatedAt,
			&project.UpdatedAt,
			&project.UserId,
		)
		helper.PanicIfError(err)
		projects = append(projects, project)
	}

	return projects
}

func (r *ProjectRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Project {
	SQL := `SELECT id, name, description, progress, confidence, trend, created_at, updated_at, user_id 
			FROM projects`

	var projects []domain.Project
	var rows *sql.Rows
	var err error

	if tx != nil {
		rows, err = tx.QueryContext(ctx, SQL)
	} else {
		rows, err = r.DB.QueryContext(ctx, SQL)
	}

	helper.PanicIfError(err)
	defer rows.Close()

	for rows.Next() {
		var project domain.Project
		err := rows.Scan(
			&project.Id,
			&project.Name,
			&project.Description,
			&project.Progress,
			&project.Confidence,
			&project.Trend,
			&project.CreatedAt,
			&project.UpdatedAt,
			&project.UserId,
		)
		helper.PanicIfError(err)
		projects = append(projects, project)
	}

	return projects
}
