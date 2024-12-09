package domain

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/DonCuervoC/banking_go_api_hex/errs"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type CustomerRepositoryDb struct {
	client *sql.DB
}

func (d CustomerRepositoryDb) FindAll() ([]Customer, error) {

	findAllSql := `SELECT customer_id, name, city, zipcode, date_of_birth, status 
					 FROM customers`

	rows, err := d.client.Query(findAllSql)

	if err != nil {
		log.Println("Error while querying customer table ", err.Error())
		return nil, err
	}

	customers := make([]Customer, 0)
	for rows.Next() {

		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.ZipCode, &c.DateOfBirth, &c.Status)
		if err != nil {
			log.Println("Error while scanning customers ", err.Error())
			return nil, err
		}
		customers = append(customers, c)
	}

	return customers, nil

}

func (d CustomerRepositoryDb) FindById(id string) (*Customer, *errs.AppError) {
	querySql := `
        SELECT customer_id, name, city, zipcode, date_of_birth, status 
        FROM customers 
        WHERE customer_id = $1` // Usamos $1 como placeholder

	// Pasamos el valor de id como argumento para QueryRow.
	row := d.client.QueryRow(querySql, id)

	var c Customer
	err := row.Scan(&c.Id, &c.Name, &c.City, &c.ZipCode, &c.DateOfBirth, &c.Status)
	if err != nil {

		if err == sql.ErrNoRows {
			// log.Println("No customer found with id:", id)
			//return nil, nil // No error, simplemente no hay datos.
			// return nil, errors.New("customer not found")
			return nil, errs.NewNotFoundError("customer not found")
		} else {
			//log.Println("Error while scanning customer:", err.Error())
			// return nil, errors.New("unexpected database error")
			return nil, errs.NewUnexpectedError("unexpected db error")
			// return nil, err
		}

	}

	return &c, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Configuración de la conexión con PostgreSQL
	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"))

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)

	}

	if err := db.Ping(); err != nil {
		log.Fatal("Error while connecting to the database: ", err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return CustomerRepositoryDb{db}

}

func ConnectToPostDB() (*gorm.DB, error) {
	// Cargar variables de entorno
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return nil, err
	}

	// Configuración de la conexión con PostgreSQL
	dbURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"))

	// Conectar a la base de datos usando GORM
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Error while connecting to the database: ", err)
		return nil, err
	}

	// Configurar el tiempo máximo de vida de la conexión
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Error getting generic database object: ", err)
		return nil, err
	}

	sqlDB.SetConnMaxLifetime(time.Minute * 5) // Configurar el tiempo máximo de vida de la conexión

	return db, nil
}
