package repository

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"github.com/google/uuid"
	"task-management/helper"
	"task-management/model/domain"
)

type ProfileRepositoryImpl struct {
	DB *sql.DB
}

func NewProfileRepository(db *sql.DB) ProfileRepository {
	return &ProfileRepositoryImpl{DB: db}
}

func (r *ProfileRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, profile domain.Profile) domain.Profile {
	if profile.Id == uuid.Nil {
		profile.Id = uuid.New()
	}
	SQL := "INSERT INTO profiles(id, user_id, full_name, email, role) VALUES($1, $2, $3, $4, $5)"
	var err error
	if tx != nil {
		_, err = tx.ExecContext(ctx, SQL,
			profile.Id,
			profile.UserId,
			strings.TrimSpace(profile.FullName),
			strings.ToLower(strings.TrimSpace(profile.Email)),
			profile.Role,
		)
	} else {
		_, err = r.DB.ExecContext(ctx, SQL,
			profile.Id,
			profile.UserId,
			strings.TrimSpace(profile.FullName),
			strings.ToLower(strings.TrimSpace(profile.Email)),
			profile.Role,
		)
	}
	helper.PanicIfError(err)
	return profile
}

func (r *ProfileRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, profile domain.Profile) domain.Profile {
	SQL := "UPDATE profiles SET user_id = $1, full_name = $2, email = $3, role = $4 WHERE id = $5"
	var err error
	if tx != nil {
		_, err = tx.ExecContext(ctx, SQL,
			profile.UserId,
			strings.TrimSpace(profile.FullName),
			strings.ToLower(strings.TrimSpace(profile.Email)),
			profile.Role,
			profile.Id,
		)
	} else {
		_, err = r.DB.ExecContext(ctx, SQL,
			profile.UserId,
			strings.TrimSpace(profile.FullName),
			strings.ToLower(strings.TrimSpace(profile.Email)),
			profile.Role,
			profile.Id,
		)
	}
	helper.PanicIfError(err)
	return profile
}

func (r *ProfileRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, profileId uuid.UUID) {
	SQL := "DELETE FROM profiles WHERE id = $1"
	var err error
	if tx != nil {
		_, err = tx.ExecContext(ctx, SQL, profileId)
	} else {
		_, err = r.DB.ExecContext(ctx, SQL, profileId)
	}
	helper.PanicIfError(err)
}

func (r *ProfileRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, profileId uuid.UUID) (domain.Profile, error) {
	SQL := "SELECT id, user_id, full_name, email, role FROM profiles WHERE id = $1"
	var profile domain.Profile
	var err error
	if tx != nil {
		err = tx.QueryRowContext(ctx, SQL, profileId).Scan(
			&profile.Id,
			&profile.UserId,
			&profile.FullName,
			&profile.Email,
			&profile.Role,
		)
	} else {
		err = r.DB.QueryRowContext(ctx, SQL, profileId).Scan(
			&profile.Id,
			&profile.UserId,
			&profile.FullName,
			&profile.Email,
			&profile.Role,
		)
	}
	if err == sql.ErrNoRows {
		return profile, errors.New("profile not found")
	}
	helper.PanicIfError(err)
	return profile, nil
}

func (r *ProfileRepositoryImpl) FindByUserId(ctx context.Context, tx *sql.Tx, userId uuid.UUID) (domain.Profile, error) {
	SQL := "SELECT id, user_id, full_name, email, role FROM profiles WHERE user_id = $1"
	var profile domain.Profile
	var err error
	if tx != nil {
		err = tx.QueryRowContext(ctx, SQL, userId).Scan(
			&profile.Id,
			&profile.UserId,
			&profile.FullName,
			&profile.Email,
			&profile.Role,
		)
	} else {
		err = r.DB.QueryRowContext(ctx, SQL, userId).Scan(
			&profile.Id,
			&profile.UserId,
			&profile.FullName,
			&profile.Email,
			&profile.Role,
		)
	}
	if err == sql.ErrNoRows {
		return profile, errors.New("profile not found")
	}
	helper.PanicIfError(err)
	return profile, nil
}

func (r *ProfileRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Profile {
	SQL := "SELECT id, user_id, full_name, email, role FROM profiles"
	var profiles []domain.Profile
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
		var profile domain.Profile
		err := rows.Scan(
			&profile.Id,
			&profile.UserId,
			&profile.FullName,
			&profile.Email,
			&profile.Role,
		)
		helper.PanicIfError(err)
		profiles = append(profiles, profile)
	}
	return profiles
}