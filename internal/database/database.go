package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func CreateTables() *sql.DB {
	user := os.Getenv("user")         //пользователь Postgres
	password := os.Getenv("password") //Пароль Postgres
	dbname := os.Getenv("dbname")     //Название базы данных
	host := os.Getenv("dbhost")       //Хост базы данных
	port := os.Getenv("dbport")       //Прт базы данных

	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, dbname))
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Failed to create migration driver: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		"postgres", driver)
	if err != nil {
		log.Fatalf("Failed to create migrate instance: %v", err)
	}

	//Поднять миграции
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to apply migrations: %v", err)
	}

	//Удалить таблицы
	// if err := m.Down(); err != nil && err != migrate.ErrNoChange {
	// 	log.Fatalf("Failed to down migrations: %v", err)
	// }

	log.Println("Migrations applied successfully!")
	return db
}
