package repository

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"github.com/google/uuid"
	"task-management/helper"
	"task-management/model/domain"
)

type UserRepositoryImpl struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{DB: db}
}

func (r *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	if user.Id == uuid.Nil {
		user.Id = uuid.New()
	}
	SQL := "INSERT INTO users(id, full_name, email, password_hash, role) VALUES($1, $2, $3, $4, $5)"
	var err error
	if tx != nil {
		_, err = tx.ExecContext(ctx, SQL,
			user.Id,
			strings.TrimSpace(user.FullName),
			strings.ToLower(strings.TrimSpace(user.Email)),
			user.PasswordHash,
			user.Role,
		)
	} else {
		_, err = r.DB.ExecContext(ctx, SQL,
			user.Id,
			strings.TrimSpace(user.FullName),
			strings.ToLower(strings.TrimSpace(user.Email)),
			user.PasswordHash,
			user.Role,
		)
	}
	helper.PanicIfError(err)
	return user
}

func (r *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	// Ambil data lama
	var oldUser domain.User
	err := r.DB.QueryRowContext(ctx, "SELECT full_name, email, password_hash, role FROM users WHERE id=$1 AND deleted_at IS NULL", user.Id).Scan(
		&oldUser.FullName,
		&oldUser.Email,
		&oldUser.PasswordHash,
		&oldUser.Role,
	)
	helper.PanicIfError(err)

	// Gunakan data lama jika field kosong
	if strings.TrimSpace(user.FullName) == "" {
		user.FullName = oldUser.FullName
	}
	if strings.TrimSpace(user.Email) == "" {
		user.Email = oldUser.Email
	}
	if strings.TrimSpace(user.PasswordHash) == "" {
		user.PasswordHash = oldUser.PasswordHash
	}
	if strings.TrimSpace(user.Role) == "" {
		user.Role = oldUser.Role
	}

	SQL := "UPDATE users SET full_name=$1, email=$2, password_hash=$3, role=$4 WHERE id=$5 AND deleted_at IS NULL"

	if tx != nil {
		_, err = tx.ExecContext(ctx, SQL,
			strings.TrimSpace(user.FullName),
			strings.ToLower(strings.TrimSpace(user.Email)),
			user.PasswordHash,
			user.Role,
			user.Id,
		)
	} else {
		_, err = r.DB.ExecContext(ctx, SQL,
			strings.TrimSpace(user.FullName),
			strings.ToLower(strings.TrimSpace(user.Email)),
			user.PasswordHash,
			user.Role,
			user.Id,
		)
	}
	helper.PanicIfError(err)

	return user
}

func (r *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, userId uuid.UUID) {
	SQL := "UPDATE users SET deleted_at=$1 WHERE id=$2"
	var err error
	if tx != nil {
		_, err = tx.ExecContext(ctx, SQL, time.Now(), userId)
	} else {
		_, err = r.DB.ExecContext(ctx, SQL, time.Now(), userId)
	}
	helper.PanicIfError(err)
}

func (r *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId uuid.UUID) (domain.User, error) {
	SQL := "SELECT id, full_name, email, password_hash, role FROM users WHERE id=$1 AND deleted_at IS NULL"
	var row *sql.Row
	if tx != nil {
		row = tx.QueryRowContext(ctx, SQL, userId)
	} else {
		row = r.DB.QueryRowContext(ctx, SQL, userId)
	}

	user := domain.User{}
	err := row.Scan(&user.Id, &user.FullName, &user.Email, &user.PasswordHash, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.User{Id: uuid.Nil}, nil
		}
		return domain.User{}, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.User, error) {
	email = strings.ToLower(strings.TrimSpace(email))
	SQL := "SELECT id, full_name, email, password_hash, role FROM users WHERE email=$1 AND deleted_at IS NULL"

	var row *sql.Row
	if tx != nil {
		row = tx.QueryRowContext(ctx, SQL, email)
	} else {
		row = r.DB.QueryRowContext(ctx, SQL, email)
	}

	user := domain.User{}
	err := row.Scan(&user.Id, &user.FullName, &user.Email, &user.PasswordHash, &user.Role)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.User{Id: uuid.Nil}, nil
		}
		return domain.User{}, err
	}
	return user, nil
}

func (r *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.User {
	SQL := "SELECT id, full_name, email, password_hash, role FROM users WHERE deleted_at IS NULL"
	var rows *sql.Rows
	var err error
	if tx != nil {
		rows, err = tx.QueryContext(ctx, SQL)
	} else {
		rows, err = r.DB.QueryContext(ctx, SQL)
	}
	helper.PanicIfError(err)
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		user := domain.User{}
		err := rows.Scan(&user.Id, &user.FullName, &user.Email, &user.PasswordHash, &user.Role)
		helper.PanicIfError(err)
		users = append(users, user)
	}
	return users
}
