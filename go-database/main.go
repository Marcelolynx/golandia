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

	patients := NewPatients("Ademir Maldonado", "67984319090", "ademiradcem@gmail.com")
	err = insertPacient(db, patients)
	if err != nil {
		panic(err)
	}

	patients.Phone = "67999005511"
	err = updatePatients(db, patients)
	if err != nil {
		panic(err)
	}

	patient, err := selectOnePatient(db, patients.ID)
	if err != nil {
		panic(err)
	}
	fmt.Println(patient.Name + " Inserted successfully!")

	patientsList, err := selectAllPatients(db)
	if err != nil {
		panic(err)
	}
	fmt.Println("pacientes cadastrados:\n")
	for _, p := range patientsList {
		fmt.Println(p.Name, p.Phone, p.Email)
	}

	err = deletePatient(db, patient.ID)
	if err != nil {
		panic(err)
	}

}

func insertPacient(db *sql.DB, pacient *Patient) error {
	stmt, err := db.Prepare("INSERT INTO patients (id, name, phone, email) VALUES ($1, $2, $3, $4)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(pacient.ID, pacient.Name, pacient.Phone, pacient.Email)
	if err != nil {
		panic(err)
	}
	fmt.Println("Product successfully inserted!")
	return nil
}

func updatePatients(db *sql.DB, patients *Patient) error {
	stmt, err := db.Prepare("UPDATE patients SET name = $1, phone = $2, email = $3 WHERE id = $4")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(patients.Name, patients.Phone, patients.Email, patients.ID)
	if err != nil {
		panic(err)
	}
	fmt.Println("Product successfully updated!")
	return nil
}

func selectOnePatient(db *sql.DB, id string) (*Patient, error) {
	stmt, err := db.Prepare("SELECT id, name, phone, email FROM patients WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var patient Patient
	err = stmt.QueryRow(id).Scan(&patient.ID, &patient.Name, &patient.Phone, &patient.Email)
	if err != nil {
		return nil, err
	}
	return &patient, nil
}

func selectAllPatients(db *sql.DB) ([]Patient, error) {
	rows, err := db.Query("SELECT id, name, phone, email FROM patients")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var patients []Patient
	for rows.Next() {
		var p Patient
		err = rows.Scan(&p.ID, &p.Name, &p.Phone, &p.Email)
		if err != nil {
			return nil, err
		}
		patients = append(patients, p)
	}
	return patients, nil
}

func deletePatient(db *sql.DB, id string) error {
	stmt, err := db.Prepare("DELETE FROM patients WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	fmt.Println("Product successfully deleted!")
	return nil
}
