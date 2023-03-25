package main

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type Patient struct {
	ID    string
	Name  string
	Phone string
	Email string
}

func NewPatients(name string, phone string, email string) *Patient {
	return &Patient{
		ID:    uuid.New().String(),
		Name:  name,
		Phone: phone,
		Email: email,
	}
}

func main() {
	connStr := "user=postgres dbname=goexpert password=mco02jgp host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	patients := NewPatients("Marcelo Oliveira", "67992322024", "marcelolynx@gmail.com")
	err = insertPacient(db, patients)
	if err != nil {
		panic(err)
	}
}

func insertPacient(db *sql.DB, pacient *Patient) error {
	stmt, err := db.Prepare("INSERT INTO patients (name, phone, email) VALUES ($1, $2, $3)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(pacient.Name, pacient.Phone, pacient.Email)
	if err != nil {
		panic(err)
	}
	fmt.Println("Product successfully inserted!")
	return nil
}
