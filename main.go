package main

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"
)

type Paciente struct {
	Nome  string `json:"nome"`
	Idade int    `json:"idade"`
	Email string `json:"email"`
	Endereco
}

type Endereco struct {
	Rua    string `json:"rua"`
	Numero int    `json:"numero"`
	Cidade string `json:"cidade"`
	Bairro string `json:"bairro"`
}
type Car struct {
	Nome  string `json:"nome"`
	Marca string `json:"marca"`
}

var cars []Car
var pacientes []Paciente

func createPaciente() {
	pacientes = append(pacientes, Paciente{"Marcelo", 46, "marcelolynx@gmail.com", Endereco{"Rua 1", 123, "São Paulo", "Vila Mariana"}})
}

func createCars() {
	cars = append(cars, Car{"Gol", "Volkswagen"})
	cars = append(cars, Car{"Uno", "Fiat"})
	cars = append(cars, Car{"Celta", "Chevrolet"})
}

func main() {
	createCars()
	createPaciente()
	e := echo.New()
	e.GET("/cars", getCars)
	e.POST("/cars", postCars)
	e.GET("/pacientes", getPacientes)
	e.POST("/pacientes", postPacientes)
	e.Logger.Fatal(e.Start(":8080"))
}

func getCars(c echo.Context) error {
	return c.JSON(http.StatusOK, cars)
}

func getPacientes(c echo.Context) error {
	return c.JSON(http.StatusOK, pacientes)
}

func postCars(c echo.Context) error {
	car := new(Car)
	if err := c.Bind(car); err != nil {
		return err
	}
	cars = append(cars, *car)
	saveCar(*car)
	return c.JSON(http.StatusCreated, car)
}

func postPacientes(c echo.Context) error {
	paciente := new(Paciente)
	if err := c.Bind(paciente); err != nil {
		return err
	}
	pacientes = append(pacientes, *paciente)
	savePaciente(*paciente)
	return c.JSON(http.StatusCreated, paciente)
}

func savePaciente(paciente Paciente) error {
	db, err := sql.Open("sqlite3", "./cars.db")
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO pacientes(nome, idade, email, rua,  cidade, bairro) values($1, $2, $3, $4, $5, $6)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(paciente.Nome, paciente.Idade, paciente.Email, paciente.Rua, paciente.Cidade, paciente.Bairro)
	if err != nil {
		return err
	}

	return nil
}

func saveCar(car Car) error {
	db, err := sql.Open("sqlite3", "./cars.db")
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO cars(nome, marca) values($1, $2)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(car.Nome, car.Marca)
	if err != nil {
		return err
	}

	return nil
}
