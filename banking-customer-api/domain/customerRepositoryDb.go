package domain

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/DonCuervoC/banking_go_api_hex/errs"
	"github.com/DonCuervoC/banking_go_api_hex/logger"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	var findAllSql string

	// Declaración de la variable rows fuera del switch para que sea accesible en todos los casos
	//var rows *sql.Rows
	var err error
	// Procesar los resultados de la consulta
	customers := make([]Customer, 0)

	// Lógica para construir la consulta SQL dependiendo del estado
	switch status {
	case "active":
		findAllSql = `SELECT customer_id, name, city, zipcode, date_of_birth, status 
						FROM customers WHERE status = TRUE`
		// rows, err = d.client.Query(findAllSql)
		err = d.client.Select(&customers, findAllSql)

	case "inactive":
		findAllSql = `SELECT customer_id, name, city, zipcode, date_of_birth, status 
						FROM customers WHERE status = FALSE`
		// rows, err = d.client.Query(findAllSql)
		err = d.client.Select(&customers, findAllSql)

	default:
		findAllSql = `SELECT customer_id, name, city, zipcode, date_of_birth, status 
						FROM customers`
		// rows, err = d.client.Query(findAllSql)
		err = d.client.Select(&customers, findAllSql)
	}

	// Comprobar si hubo un error al ejecutar la consulta SQL
	if err != nil {
		// log.Println("Error while querying customer table: ", err.Error())
		logger.Error("Error while querying customer table: " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected database error")
	}

	// Retornar los clientes obtenidos
	return customers, nil
}

func (d CustomerRepositoryDb) FindById(id string) (*Customer, *errs.AppError) {
	querySql := `
        SELECT customer_id, name, city, zipcode, date_of_birth, status 
        FROM customers 
        WHERE customer_id = $1` // Usamos $1 como placeholder

	var c Customer
	err := d.client.Get(&c, querySql, id)

	if err != nil {

		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		} else {
			return nil, errs.NewUnexpectedError("unexpected db error")
		}
	}
	return &c, nil
}

func (d CustomerRepositoryDb) CustomerExists(id string) (bool, *errs.AppError) {
	querySql := `
        SELECT COUNT(*) 
        FROM customers 
        WHERE customer_id = $1;
    `

	var count int
	err := d.client.Get(&count, querySql, id)

	if err != nil {
		logger.Error("Error while checking if customer exists: " + err.Error())
		return false, errs.NewUnexpectedError("unexpected database error")
	}

	// Si el conteo es mayor a 0, el cliente existe
	return count > 0, nil
}

func NewCustomerRepositoryDb(dbClient *sqlx.DB) CustomerRepositoryDb {
	return CustomerRepositoryDb{dbClient}
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
