package app

import (
	"database/sql"
	"fmt"
	"time"

	"task-management/helper"
	"task-management/model/domain"

	_ "github.com/jackc/pgx/v5/stdlib"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDB() *sql.DB {
	// GORM setup for auto-migration
	dsn := "postgres://postgres:postgres@localhost:5435/task-management-sarana?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	helper.PanicIfError(err)

	// Auto migrate all models
	// err = db.AutoMigrate(
	// 	&domain.User{},
	// 	&domain.Task{},
	// 	&domain.Project{},
	// 	&domain.Profile{},
	// 	&domain.RefreshToken{},
	// )
	if err != nil {
		panic(fmt.Sprintf("❌ Auto-migration failed: %v", err))
	}

	// Add specific columns if they don't exist
	migrator := db.Migrator()
	
	// Add Progress column if it doesn't exist
	if !migrator.HasColumn(&domain.Task{}, "progress") {
		err = migrator.AddColumn(&domain.Task{}, "progress")
		if err != nil {
			fmt.Printf("⚠️ Failed to add progress column: %v\n", err)
		} else {
			fmt.Println("✅ Added progress column to tasks table")
		}
	}
	
	// Add ContinueTomorrow column if it doesn't exist
	if !migrator.HasColumn(&domain.Task{}, "continue_tomorrow") {
		err = migrator.AddColumn(&domain.Task{}, "continue_tomorrow")
		if err != nil {
			fmt.Printf("⚠️ Failed to add continue_tomorrow column: %v\n", err)
		} else {
			fmt.Println("✅ Added continue_tomorrow column to tasks table")
		}
	}

	fmt.Println("✅ Auto-migration completed successfully!")

	// Get underlying sql.DB for compatibility with existing code
	sqlDB, err := db.DB()
	helper.PanicIfError(err)

	// Test connection
	err = sqlDB.Ping()
	if err != nil {
		panic(fmt.Sprintf("❌ Gagal koneksi ke database: %v", err))
	}
	fmt.Println("✅ Database berhasil terkoneksi!")

	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxLifetime(60 * time.Minute)
	sqlDB.SetConnMaxIdleTime(10 * time.Minute)

	// db.Migrator().AddColumn(&domain.Task{}, "progress")
	// db.Migrator().AddColumn(&domain.Task{}, "continue_tomorrow")
	// db.AutoMigrate()
	return sqlDB
}
