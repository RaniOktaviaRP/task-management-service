package app

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"task-management/helper"
)

func NewDB() *sql.DB {
	db, err := sql.Open("pgx", "postgres://postgres:rani2510@localhost:5432/Management?sslmode=disable")
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
