package app

import (
	"database/sql"
	"fmt"
	"time"

	"task-management/helper"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewDB() *sql.DB {
	db, err := sql.Open("pgx", "postgres://postgres:postgres@localhost:5435/task-management-sarana?sslmode=disable")
	helper.PanicIfError(err)

	// cek koneksi
	err = db.Ping()
	if err != nil {
		panic(fmt.Sprintf("❌ Gagal koneksi ke database: %v", err))
	}
	fmt.Println("✅ Database berhasil terkoneksi!")

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
