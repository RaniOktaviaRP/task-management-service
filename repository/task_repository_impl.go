package repository

import (
	"context"
	"database/sql"
	"errors"
	"task-management/model/domain"
	"time"

	"github.com/google/uuid"
)

type TaskRepositoryImpl struct {
	DB *sql.DB
}

func NewTaskRepository(db *sql.DB) TaskRepository {
	return &TaskRepositoryImpl{
		DB: db,
	}
}

func (repository *TaskRepositoryImpl) Save(ctx context.Context, task domain.Task) (domain.Task, error) {
	query := `INSERT INTO tasks (id, project_id, title, status, priority, effort, difficulty_level, deliverable, bottleneck, progress, continue_tomorrow, created_at, updated_at)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)`

	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	_, err := repository.DB.ExecContext(ctx, query,
		task.Id, task.ProjectId, task.Title, task.Status, task.Priority,
		task.Effort, task.DifficultyLevel, task.Deliverable, task.Bottleneck,
		task.Progress, task.ContinueTomorrow,
		task.CreatedAt, task.UpdatedAt)

	if err != nil {
		return task, err
	}
	return task, nil
}

func (repository *TaskRepositoryImpl) Update(ctx context.Context, task domain.Task) (domain.Task, error) {
	query := `UPDATE tasks SET 
		project_id = $1, title = $2, status = $3, priority = $4,
		effort = $5, difficulty_level = $6, deliverable = $7, bottleneck = $8,
		progress = $9, continue_tomorrow = $10, updated_at = $11
		WHERE id = $12`

	task.UpdatedAt = time.Now()

	result, err := repository.DB.ExecContext(ctx, query,
		task.ProjectId, task.Title, task.Status, task.Priority,
		task.Effort, task.DifficultyLevel, task.Deliverable, task.Bottleneck,
		task.Progress, task.ContinueTomorrow, task.UpdatedAt, task.Id)

	if err != nil {
		return task, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return task, err
	}

	if rowsAffected == 0 {
		return task, errors.New("task not found")
	}

	return task, nil
}

func (repository *TaskRepositoryImpl) Delete(ctx context.Context, taskId uuid.UUID) error {
	query := `DELETE FROM tasks WHERE id = $1`

	result, err := repository.DB.ExecContext(ctx, query, taskId)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("task not found")
	}

	return nil
}

func (repository *TaskRepositoryImpl) FindById(ctx context.Context, taskId uuid.UUID) (domain.Task, error) {
	query := `SELECT id, project_id, title, status, priority, effort, difficulty_level, deliverable, bottleneck, progress, continue_tomorrow, created_at, updated_at
		FROM tasks WHERE id = $1`

	var task domain.Task
	var progress sql.NullString
	var continueTomorrow sql.NullBool
	
	err := repository.DB.QueryRowContext(ctx, query, taskId).Scan(
		&task.Id, &task.ProjectId, &task.Title, &task.Status, &task.Priority,
		&task.Effort, &task.DifficultyLevel, &task.Deliverable, &task.Bottleneck,
		&progress, &continueTomorrow,
		&task.CreatedAt, &task.UpdatedAt)

	if err == sql.ErrNoRows {
		return task, errors.New("task not found")
	}

	if err != nil {
		return task, err
	}

	// Handle NULL values
	if progress.Valid {
		task.Progress = progress.String
	} else {
		task.Progress = ""
	}
	
	if continueTomorrow.Valid {
		task.ContinueTomorrow = continueTomorrow.Bool
	} else {
		task.ContinueTomorrow = false
	}

	return task, nil
}

func (repository *TaskRepositoryImpl) FindByProjectId(ctx context.Context, projectId uuid.UUID) ([]domain.Task, error) {
	query := `SELECT id, project_id, title, status, priority, effort, difficulty_level, deliverable, bottleneck, progress, continue_tomorrow, created_at, updated_at
		FROM tasks WHERE project_id = $1`

	rows, err := repository.DB.QueryContext(ctx, query, projectId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []domain.Task
	for rows.Next() {
		var task domain.Task
		var progress sql.NullString
		var continueTomorrow sql.NullBool
		
		err := rows.Scan(
			&task.Id, &task.ProjectId, &task.Title, &task.Status, &task.Priority,
			&task.Effort, &task.DifficultyLevel, &task.Deliverable, &task.Bottleneck,
			&progress, &continueTomorrow,
			&task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return nil, err
		}
		
		// Handle NULL values
		if progress.Valid {
			task.Progress = progress.String
		} else {
			task.Progress = ""
		}
		
		if continueTomorrow.Valid {
			task.ContinueTomorrow = continueTomorrow.Bool
		} else {
			task.ContinueTomorrow = false
		}
		
		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}

func (repository *TaskRepositoryImpl) FindAll(ctx context.Context) ([]domain.Task, error) {
	query := `SELECT id, project_id, title, status, priority, effort, difficulty_level, deliverable, bottleneck, progress, continue_tomorrow, created_at, updated_at
		FROM tasks`

	rows, err := repository.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []domain.Task
	for rows.Next() {
		var task domain.Task
		var progress sql.NullString
		var continueTomorrow sql.NullBool
		
		err := rows.Scan(
			&task.Id, &task.ProjectId, &task.Title, &task.Status, &task.Priority,
			&task.Effort, &task.DifficultyLevel, &task.Deliverable, &task.Bottleneck,
			&progress, &continueTomorrow,
			&task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			return nil, err
		}
		
		// Handle NULL values
		if progress.Valid {
			task.Progress = progress.String
		} else {
			task.Progress = ""
		}
		
		if continueTomorrow.Valid {
			task.ContinueTomorrow = continueTomorrow.Bool
		} else {
			task.ContinueTomorrow = false
		}
		
		tasks = append(tasks, task)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}